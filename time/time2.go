package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mlaoji/ygo/lib"
)

// GetDateUnit get date unit format 20180102
func GetDateUnit(t time.Time) (int, error) {
	str := t.Format("20060102")
	return strconv.Atoi(str)
}

func GetDateUnitStr(str string) (int, error) {
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(t.Format("20060102"))
}

// YesterdayDate get date yesterday
func YesterdayDate(date int) int {
	t, _ := time.ParseInLocation("20060102", lib.AsString(date), time.Local)
	tY, _ := GetDateUnit(t.AddDate(0, 0, -1))
	return tY
}

func main() {
	tB := time.Now()
	time.Sleep(time.Second * 3)
	tE := time.Now()
	fmt.Println("escape:", tE.Sub(tB))
	t := FormatIntDate(20180332)
	fmt.Println(t)
	return
	d := GetDate("2018-01-02-12:02:03")
	fmt.Println(d)
	return

	y := YesterdayDate(20180101)
	fmt.Println(y)
	return

	fmt.Println(lib.UTCDate())
	fmt.Println(lib.UTCDateTime())
	fmt.Println(lib.ToInt(lib.UTCDate()))
	fmt.Println(GetDateUnitStr(lib.UTCDate()))
	return

	val := lib.UTCDateTime()
	fmt.Println(val)
	return
	l, _ := time.LoadLocation("")
	fmt.Println(l)
	tUTC, _ := time.ParseInLocation("2006-01-02 15:04:05", val, l)
	fmt.Println(lib.DateTime(int(tUTC.Unix())))

	fmt.Println(lib.DateTime(int(time.Now().Unix())))
	fmt.Println(lib.DateTime(int(time.Now().UTC().Unix())))

	fmt.Println(lib.UTCDateTime())

	fmt.Println(time.Now().UTC())
	fmt.Println(time.Now())

	fmt.Println(time.Now().UTC().Unix())
	fmt.Println(time.Now().Unix())
	return

	locla := time.Now()
	fmt.Println(locla)

	locla2 := time.Now().UTC()
	fmt.Println(locla2)
	return
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

// GetDate str 2006-01-02 15:04:05 return 2006-01-02
func GetDate(str string) string {
	s := strings.Split(str, " ")
	if len(s) > 0 {
		return s[0]
	}
	return str
}

func FormatIntDate(date int) string {
	s := strconv.Itoa(date)
	y := string([]byte(s)[:4])
	m := string([]byte(s)[4:6])
	d := string([]byte(s)[6:8])
	return fmt.Sprintf("%v-%v-%v", y, m, d)
}
