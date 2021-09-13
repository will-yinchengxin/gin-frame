package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经弃用或不存在, 请更换code", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Mag() string {
	return e.msg
}

func (e *Error) Magf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, i2 := range details {
		newError.details = append(newError.details, i2)
	}
	return &newError
}

func (e *Error) StatusCode() int {
	switch e.code {
	case Success.code:
		return http.StatusOK
	case ServerError.code:
		return http.StatusInternalServerError
	case InvalidParameter.code:
		return http.StatusBadRequest
	case AuthServerGetDataFailed.code:
		// 可以使用fallthrough强制执行后面的case代码
		fallthrough
	case TokenTypeError.code:
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}