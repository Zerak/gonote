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

	getTimestampOfTime()

	tick2()

	// 5天前的时间
	t5t := time.Now().Add(-time.Duration(time.Hour * 24 * time.Duration(5)))
	fmt.Println("5 day ago:", t5t)
	fmt.Println("5 day ago timestamp:", t5t.Unix())

	// 取当天零点
	dz := getDayTodayZero(time.Now().Unix())
	dzdate := timestampToDate(dz)
	fmt.Printf("date:%v getDayTodayZero:%v\n", dzdate, dz)

	dm := getDayYesterdayMidnight(time.Now().Unix())
	dmdate := timestampToDate(dm)
	fmt.Printf("date:%v getDayYesterdayMidnight:%v\n", dmdate, dm)

	dtm := getDayTodayMidnight(time.Now().Unix())
	dtmdate := timestampToDate(dtm)
	fmt.Printf("date:%v getDayTodayMidnight:%v\n", dtmdate, dtm)

	now := time.Now().Add(-time.Hour * 24 * 1)
	da := getDayYesterdayMidnight(now.Unix())
	fmt.Printf("getDayYesterdayMidnight :%v\n", da)
	dat := timestampToDate(da)
	fmt.Printf("date:%v\n", dat)

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

func tick2() {
	getMicro := func(n int64) (re int) {
		str := strconv.Itoa(int(n))
		fmt.Printf("str:%s\n", str)
		if len(str) > 3 {
			strN := str[len(str)-3:]
			re, _ = strconv.Atoi(strN)
		} else {
			re = int(n)
		}
		return re
	}
	now := time.Now()
	h := now.Hour()
	m := now.Minute()
	s := now.Second()
	mi := now.UnixNano() / (1000 * 1000)
	mi2 := getMicro(mi)
	fmt.Printf("%02d:%02d:%02d:%03d nano:%d mi2:%03d\n", h, m, s, mi, now.UnixNano(), mi2)
	fmt.Println("tick2")
	key := 10
	go func() {
		for {
			<-time.After(time.Second)
			key -= 1
			if key == 0 {
				key = 20
			}
		}
	}()
	for {
		//select {
		//case <-time.After(time.Minute * 15):
		//	fmt.Println("tick 15 minute")
		//case <-time.After(time.Minute * 5):
		//	fmt.Println("tick 5 minute")
		//case <-time.After(time.Minute * 1):
		//	fmt.Println("tick 1 minute")
		//case <-time.After(time.Second * 6):
		//	fmt.Println("tick 6 second")
		////default:
		////	time.Sleep(time.Microsecond)
		//}
		if key >= 1 {
			//fmt.Println("key 1...")
			select {
			case <-time.After(time.Second * 15):
				fmt.Println("tick 1 second key:", key)
			default:
			}
			//fmt.Println("key 1")
		}

		if key > 10 {
			//fmt.Println("key 10...")
			select {
			case <-time.After(time.Second * 5):
				fmt.Println("tick 2 second key:", key)
			default:
			}
			//fmt.Println("key 10")
		}

		time.Sleep(time.Microsecond)
	}
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
	date = tm.Format("20060102")

	//fmt.Printf("str date:%v\n", date)
	return date
}

func getDayTodayZero(timestamp int64) int64 {
	tm := time.Unix(timestamp, 0)
	str := tm.Format("20060102")

	t, _ := time.Parse("20060102", str)
	t = t.Add((-time.Hour * 8))

	return t.Unix()
}

func getDayYesterdayMidnight(timestamp int64) int64 {
	tm := time.Unix(timestamp, 0)
	str := tm.Format("20060102")

	t, _ := time.Parse("20060102", str)
	t = t.Add((-time.Hour * 8) - time.Second)

	return t.Unix()
}

func getDayTodayMidnight(timestamp int64) int64 {
	tm := time.Unix(timestamp, 0)
	str := tm.Format("20060102")

	t, _ := time.Parse("20060102", str)
	t = t.Add(time.Hour*16 - time.Second)

	return t.Unix()
}

func getTimestampOfTime() {
	date := time.Now().Format("20060102")
	fmt.Printf("date:%v\n", date)
	t, err := time.ParseInLocation("20060102", date, time.Local)
	//t, err := time.Parse("20060102", date)
	if err != nil {
		fmt.Printf("time parse err:%v\n", err)
	}
	t = t.Add(time.Hour * 9)
	fmt.Printf("time stamp:%v\n",t.Unix())
}
