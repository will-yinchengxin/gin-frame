package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Level int8

type Fileds map[string]interface{}

const (
	LevelDebug Level = iota // 0
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	default:
		return ""
	}
}
//----------------------------日志部分方法--------------------------------
type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fileds
	callers   []string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{
		newLogger: l,
	}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

// 设置日志公共字段
func (l *Logger) WithFields(f Fileds) *Logger {
	l1 := l.clone()
	if l1.fields == nil {
		l1.fields = make(Fileds)
	}
	for i, i2 := range f {
		l1.fields[i] = i2
	}
	return l1
}

// 设置日志上下文属性
func (l *Logger) WithContext(ctx context.Context) *Logger {
	l1 := l.clone()
	l1.ctx = ctx
	return l1
}

// 设置当前某一层调用栈的信息(程序计数器, 文件信息和行号)
func (l *Logger) WithCaller(skip int) *Logger {
	l1 := l.clone()
	// 函数把当前go程调用栈上的调用栈标识符填入切片pc中，返回写入到pc中的项数。实参skip为开始在pc中记录之前所要跳过的栈帧数，
	// 0表示Callers自身的调用栈，1表示Callers所在的调用栈。返回写入p的项数。
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		// FuncForPC返回一个表示调用栈标识符pc对应的调用栈的*Func；如果该调用栈标识符没有对应的调用栈，
		// 函数会返回nil。每一个调用栈必然是对某个函数的调用
		f := runtime.FuncForPC(pc)
		l1.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return l1
}

// 设置当前的整个调用栈信息
func (l *Logger) WithCallerFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	// 形同 for i:=0,i<10;i++, more是布尔值，表示有没有下一帧
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}
	l1 := l.clone()
	l1.callers = callers
	return l1
}

// 日志内容格式化
func (l *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	data := make(Fileds, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for i, i2 := range l.fields {
			if _, ok := data[i]; !ok {
				data[i] = i2
			}
		}
	}
	return data
}

// 日志输出动作
func (l *Logger) Output(level Level, message string) {
	body, _ := json.Marshal(l.JSONFormat(level, message))
	content := string(body)
	switch level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

// --------------------------------日志分级输出--------------------------------
/*
var test1 = []string{"will", "age"}
var test2 = map[string]interface{}{"name":"will", "age":12}
fmt.Println(fmt.Sprint(test1, test2)) // [will age] map[age:12 name:will]
*/
func (l *Logger) Debug(v...interface{})  {
	l.Output(LevelDebug, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{})  {
	l.Output(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v...interface{})  {
	l.Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{})  {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v...interface{})  {
	l.Output(LevelWarn, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{})  {
	l.Output(LevelWarn, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v...interface{})  {
	l.Output(LevelError, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{})  {
	l.Output(LevelError, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v...interface{})  {
	l.Output(LevelFatal, fmt.Sprint(v...))
}

func (l *Logger) Fatalf(format string, v ...interface{})  {
	l.Output(LevelFatal, fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(v...interface{})  {
	l.Output(LevelPanic, fmt.Sprint(v...))
}

func (l *Logger) Panicf(format string, v ...interface{})  {
	l.Output(LevelPanic, fmt.Sprintf(format, v...))
}
