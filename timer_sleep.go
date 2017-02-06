package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(int(time.Second))
	forNum := 1000 * 1000 * 1000 // us ms s
	test1(10000, forNum)
	test1(1000, forNum)
	test1(1, forNum)

}

func test1(sleepTime time.Duration, num int) {
	i := 0
	t := time.Now()
	for {
		i++
		if i >= num {
			fmt.Println("forNum:", num, " sleep:", sleepTime, " use:", time.Now().Sub(t).String())
			break
		}
		time.Sleep(sleepTime)
	}
}
