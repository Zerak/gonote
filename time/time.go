package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("after...")
	//<-time.After(time.Second * 2)
	//fmt.Println("after time...")

	y, m, d := time.Now().Date()
	fmt.Printf("year:%v, month:%v, day:%v\n", y, m, d)

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
	//time.Sleep(time.Hour * 1)
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
