package main

import (
	"fmt"
	"time"
)

func main() {
	//now := time.Now().AddDate(0, 0, -30)
	now := time.Now().Add(-(time.Hour * 24 * 30))
	fmt.Println(now, "stamp :", now.Unix())

	GetDayTodayZero(time.Now().Unix())

	GetDayTodayMidnight(time.Now().Unix())

	GetDayYesterdayMidnight(time.Now().Unix())
}

// 获取当天零点时间戳 传入20060102日期的时间戳,返回 20060102:00:00:00 的时间戳
func GetDayTodayZero(timestamp int64) int64 {
	tm := time.Unix(timestamp, 0)
	str := tm.Format("20060102")

	t, _ := time.ParseInLocation("20060102", str, time.Local)

	fmt.Printf("t:%v\n", t)
	return t.Unix()
}

// 获取当天午夜时间戳 传入20060102日期的时间戳,返回 20060102:23:59:59 的时间戳
func GetDayTodayMidnight(timestamp int64) int64 {
	tm := time.Unix(timestamp, 0)
	str := tm.Format("20060102")

	t, _ := time.ParseInLocation("20060102", str, time.Local)
	t = t.Add(time.Hour*24 - time.Second)

	fmt.Printf("t:%v\n", t)
	return t.Unix()
}

// 获取昨天零点的时间戳 传入20060102日期的时间戳,返回 20060101:23:59:59 的时间戳
func GetDayYesterdayMidnight(timestamp int64) int64 {
	tm := time.Unix(timestamp, 0)
	str := tm.Format("20060102")

	t, _ := time.ParseInLocation("20060102", str, time.Local)
	t = t.Add(-time.Second)

	fmt.Printf("t:%v\n", t)
	return t.Unix()
}
