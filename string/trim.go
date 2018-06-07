package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var uids []int
	for i := 1; i < 201; i++ {
		uids = append(uids, i)
	}

	//trim1(uids)
	trim2(uids)

	method := "/passport/getuserinfo 是/"
	s := strings.Trim(method, "/")
	fmt.Println(s)

	st := strings.Title(s)
	fmt.Println(st)
}

func trim1(uids []int) {
	var des [][]string
	var uidStr []string
	for k, v := range uids {
		if k+1 > 10000 {
			break
		}

		uidStr = append(uidStr, strconv.Itoa(v))
		if (k+1)%100 == 0 && k > 0 {
			des = append(des, uidStr)
			uidStr = []string{}
		}
	}
	if len(uidStr) > 0 {
		des = append(des, uidStr)
	}
	for _, v := range des {
		fmt.Println(v)
	}
}

func trim2(uids []int) {
	const (
		limitSize = 100
		pagePer   = 10
	)
	if len(uids) > limitSize {
		uids = uids[0:limitSize]
	}
	var des [][]int
	page := len(uids) / pagePer
	for i := 0; i < page; i++ {
		s := i * pagePer
		e := (i + 1) * pagePer
		des = append(des, uids[s:e])
	}
	left := len(uids) % pagePer
	if left > 0 {
		s := page * pagePer
		e := len(uids)
		des = append(des, uids[s:e])
	}
	for _, v := range des {
		fmt.Println(v)
	}
}
