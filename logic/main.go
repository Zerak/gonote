package main

import (
	"fmt"
	"strings"
)

func Test(source string) {
	switch source {
	case "wx":
		fallthrough
	case "qq":
		fallthrough
	case "sina":
		fmt.Println("wx,qq,sina:", source)

	case "official":
		fallthrough
	case "mobile":
		fallthrough
	case "robot":
		fmt.Println("offi,mobile,robot:", source)
	}
}

func main() {
	// test code1
	//Test("wx")
	//fmt.Println(fmt.Sprint("tt", "aa"))

	// test code2
	vids := []int{1, 2, 3, 4, 5, 6}
	if len(vids) > 5 {
		vids = vids[0:5]
	}
	fmt.Println(vids)

	// test code3
	a := []int{1, 2, 3}
	if true {
		fmt.Println(a)
		a, err := getA(a)
		fmt.Println(a, err)
		fmt.Println("end")
	}

	// test code4
	str := "update video_tb set status = 1 and des = 'right' and"
	str = strings.TrimRight(str, "and")
	fmt.Println(str)
}

func getA(a []int) (b []int, err error) {
	b = a
	b = append(b, []int{4, 5, 6}...)
	return
}
