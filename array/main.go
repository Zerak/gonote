package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wait := sync.WaitGroup{}
	num := runtime.NumCPU()
	fmt.Println("the numer of cpu core::", runtime.NumCPU())
	f := func() {
		for {
		}
		wait.Done()
	}
	for i := 0; i < num; i++ {
		wait.Add(1)
		go f()
	}

	//wait for exit
	wait.Wait()

	a := 10
	switch a {
	case 10:
		fmt.Println("case10: ", a)
	case 20:
		fmt.Println("case10: ", a)
	}
	return
	robots := make([][]int, 50)

	for i := 0; i < 50; i++ {
		robots[i][i] = i
		robots[i][i] = i * 20
	}
	fmt.Println(robots)
}
