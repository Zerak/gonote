package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//fmt.Println("after...")
	//<-time.After(time.Second * 2)
	//fmt.Println("after time...")

	// 取当天零点
	fmt.Printf("getDayTodayZero:%v\n",getDayTodayZero(time.Now().Unix()))
	fmt.Printf("getDayYesterdayZero:%v\n",getDayYesterdayZero(time.Now().Unix()))

	timestampToDate(time.Now().Unix())

	y, m, d := time.Now().Date()
	Date := fmt.Sprintf("%04d%02d%02d%02d%02d%02d", y, m, d, time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	fmt.Println("date:", Date)
	fmt.Printf("year:%v, month:%v, day:%v\n", y, m, d)

	date, _ := strconv.Atoi(fmt.Sprintf("%04d%02d%02d%02d%02d%02d", y, m, d, time.Now().Hour(), time.Now().Minute(), time.Now().Second()))
	fmt.Println("date int:", date)

	dura, _ := time.ParseDuration("-24h")
	fmt.Printf("dura:%v\n", dura)

	tt := time.Now().Add(-(time.Hour * 24 * 5))
	fmt.Printf("1days:%v\n5days:%v\n5daysNano:%v\n", time.Now(), tt, tt.UnixNano())
	//sub := tt.Sub(time.Now())
	sub := time.Now().Sub(tt)
	fmt.Printf("sub:%v\n", sub)

	fmt.Println(tt.UnixNano())

	startTimer(func() {
		fmt.Println(".")
	})

}

func startTimer(f func()) {
	// 立刻执行
	f()
	now := time.Now()
	y, m, d := now.Date()
	nextTime := time.Date(y, m, d, 15, 40, 0, 0, now.Location())
	if now.Unix() >= nextTime.Unix() {
		nextTime = nextTime.Add(24 * time.Hour)
	}
	// 等待下次执行点
	<-time.After(nextTime.Sub(now))
	f()
	// 以后每个一天执行一次
	for range time.Tick(24 * time.Hour) {
		f()
	}
}

func timestampToDate(stamp int64) (date string) {
	tm := time.Unix(stamp, 0)
	str := tm.Format("20060102")

	fmt.Printf("str date:%v\n", str)
	return date
}

func getDayTodayZero(timestamp int64) int64 {
	tm := time.Unix(timestamp, 0)
	str := tm.Format("20060102")

	t, _ := time.Parse("20060102", str)
	t = t.Add((-time.Hour * 8))

	return  t.Unix()
}

func getDayYesterdayZero(timestamp int64) int64{
	tm := time.Unix(timestamp, 0)
	str := tm.Format("20060102")

	t, _ := time.Parse("20060102", str)
	t = t.Add((-time.Hour * 8) - time.Second)

	return  t.Unix()
}