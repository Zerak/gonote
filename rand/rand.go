package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//rand.Seed(time.Now().UnixNano())

	idx := 0
	for i := 0; i < 1000; i++ {
		rand.Seed(1234567)
		fmt.Println(rand.Int63(), " ", rand.Int(), " ", 1+rand.Intn(5))
		idx = i
		break
	}
	fmt.Println(idx)
	return

	time.AfterFunc(time.Second*5, func() {
		fmt.Printf("tick\n")
	})

	for {
		select {
		case <-time.After(time.Second):
			//fmt.Printf("rand:%v\n", rand.Intn(120-30)+30)
			fmt.Printf("rand:%v\n", rand.Intn(3-0)+0)
		}
	}
	inviteCode()
}

func inviteCode() {
	for i := 10200000; i < 10400000; i++ {
		fmt.Printf("rand:%v\n", rand.Intn(120-30)+30)
		code := strconv.FormatInt(int64(i), 16)
		if code == "9ca418" {
			fmt.Printf("code:%v uid:%v\n", code, i)
		}
	}
}
