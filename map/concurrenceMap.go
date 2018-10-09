package main

import (
	"fmt"
	"sync"
)

var m = make(map[int]int)

func consume(i int, wg *sync.WaitGroup) {
	for k, v := range m {
		fmt.Println("routin:", i, k, v)
	}
	wg.Done()
}

func main() {
	l := 1000000
	for i := 0; i < l; i++ {
		m[i] = i * 10
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go consume(i, &wg)
	}

	wg.Wait()
}
