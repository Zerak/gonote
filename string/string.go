package main

import (
	"errors"
	"fmt"
	"hash/fnv"
	"math"
	"strconv"
	"strings"
	"time"
)

var (
	argError = errors.New("参数错误")
)

func Hash(s string) uint64 {
	h := fnv.New64()
	h.Write([]byte(s))
	return h.Sum64()
}

func getDate(date int) string {
	d := strconv.Itoa(date)
	d = string([]byte(d)[2:])
	return d
}

func main() {
	fmt.Println(getDate(20180102))
	return
	a := 0.1
	b := 0.2
	c := a + b
	fmt.Println(c)

	score := 376
	weight := 22
	val := float64(score) + float64(score)*float64(weight)/100.0
	fmt.Println(val)
	fmt.Println(int(val))
	fmt.Println(math.Trunc(val))

	val2 := score + score*weight/100
	fmt.Println("v21:", val2)
	fmt.Println("v22:", val2)
	return

	uuid := time.Now().UnixNano()
	s := fmt.Sprintf("%v", uuid)
	fmt.Println(s)
	return

	android := "ANDROID"
	fmt.Println(strings.Split(android, ","))
	return
	//did := "C7DDC9A6-3448-4945-8CB4-02DF983049D2"
	did := "f9ba2ecad7dbf3596a990bd5289e0103"
	key := Hash(did)
	fmt.Println("key:", key, " hash:", key%10)
	return
	strTest := "luuLnNfK3mYaOget_fJYX5XYP8uF?vframe/jpg/offset/0"
	strArr := strings.Split(strTest, ",")
	fmt.Println("len:", len(strArr), " data:", strArr)
	return
	id := "123412345678"
	//id := "123500000000"
	//id := strings.Fields(idStr)
	//for k, v := range id {
	//	fmt.Println(k, v)
	//}
	class, _ := strconv.Atoi(id[0:4])
	tag, _ := strconv.Atoi(id[4:12])
	fmt.Println("class:", class, " tag:", tag)

	n := (class + 1) * 1000 * 1000 * 100

	idInt, _ := strconv.Atoi(id)
	t := idInt + 1
	fmt.Println("next class:", n, " tag:", t)

	// t[0:4]
	nC := t / (1000 * 1000 * 100)
	fmt.Println("nc:", nC)
	if nC != class {
		fmt.Println("err tid")
	}
	return

	str := ",1,2,3,3,4,"
	str = strings.TrimRight(str, ",")
	fmt.Println("trim->", str)
	str = strings.TrimLeft(str, ",")
	fmt.Println("trim->", str)
	fmt.Printf("arg:%+v\n", argError)
	for i := 10000000; i < 10000099; i++ {
		fmt.Println(encodeUserId(i))
	}
}

func encodeUserId(uid int) string {
	base := 10000000
	if uid < base {
		return "100***00"
	}
	tail := uid % base
	return fmt.Sprintf("100***%02d %v", tail, tail)
}
