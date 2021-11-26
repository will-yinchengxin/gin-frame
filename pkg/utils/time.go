package utils

import (
	"fmt"
	"time"
)

// 获取指定时间戳 0 点和 23:59 时间戳
func ParseDate(timeInt int64) (start, end int64) {
	timerNow := time.Unix(timeInt, 0).Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timerNow)
	timeStart := t.Unix()
	timeEnd := timeStart + 86400000 - 1
	return timeStart, timeEnd
}

// 获取当天零点的时间戳
func TodayZero()  {
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println("timeStr:", timeStr)
	t, _ := time.Parse("2006-01-02", timeStr)
	timeNumber := t.Unix()
	fmt.Println("timeNumber:", timeNumber)
}

// 计算两个时间戳之间的所有时间戳
// 请求: time.StatisticsDate(1630290706, 1636945034)
// [1636905600 1636819200 1636732800 1636646400 1636560000]
func StatisticsDate(startTime int64, endTime int64) (dates []int64) {
	dates = []int64{}
	tEnd := time.Unix(endTime, 0)
	tEnd = time.Date(tEnd.Year(), tEnd.Month(), tEnd.Day(), 0, 0, 0, 0, time.Local)

	tStart := time.Unix(startTime, 0)
	tStart = time.Date(tStart.Year(), tStart.Month(), tStart.Day(), 0, 0, 0, 0, time.Local)

	for {
		dates = append(dates, tEnd.Unix())
		tEnd = tEnd.AddDate(0, 0, -1)
		if tEnd.Unix() < tStart.Unix() {
			break
		}

	}
	return
}

// 计算两个时间间的时间间隔 时间格式形如: 2006/01/02
func GetTimeArr(start, end string) int64{
	timeLayout  := "2006-01-02"
	loc, _ := time.LoadLocation("Local")
	// 转成时间戳
	startUnix, _ := time.ParseInLocation(timeLayout,  start,  loc)
	endUnix, _ := time.ParseInLocation(timeLayout,  end,  loc)
	startTime := startUnix.Unix()
	endTime := endUnix.Unix()
	// 求相差天数
	date :=	(endTime - startTime) / 86400
	return date
}