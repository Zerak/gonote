package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//m := new(sync.Map)
	//for i := 0; i < 100; i++ {
	//	m.Store(fmt.Sprintf("%v", i), i)
	//}
	//
	//m.Range(func(key, value interface{}) bool {
	//	fmt.Println("Key:", key)
	//	fmt.Println("Val:", value)
	//	return true
	//})

	f, err := os.OpenFile("/Users/zerak/Project/goPro/src/gonote/log/more.log", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}

	user := make(map[int]int)
	r := bufio.NewReader(f)
	l, _, err := r.ReadLine()
	i := 0
	for err == nil {
		l, _, err = r.ReadLine()
		//fmt.Printf("line:%v\n", string(l))

		s := strings.Split(string(l), " ")
		if len(s) > 3 {
			i++
			if i > 12163 {
				fmt.Println(i)
			}
			//fmt.Printf("i:%v inviter:%v beInviter:%v amount:%v\n", i, s[0], s[1], s[2])
			inviter, _ := strconv.Atoi(s[0])
			amount, _ := strconv.Atoi(s[2])
			user[inviter] += amount
		}
	}
	total := 0
	for k, v := range user {
		fmt.Printf("%8d %v\n", k, v)
		total += v
	}
	fmt.Printf("len:%v a:%v", len(user), total)
}
