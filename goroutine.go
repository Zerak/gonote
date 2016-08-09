package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exit chan byte

func init() {
	exit = make(chan byte, 1)
}

func main() {
	wg.Add(1)
	go testGoroutine()
	wg.Wait()
}

func testGoroutine() {
	tick := time.NewTicker(2 * time.Second)
	tickc := tick.C
	sum := 0
	quit := false
	for {
		select {
		case <-tickc:
			fmt.Printf(".")
			sum += 5
			if sum > 10 {
				exit <- 1
			}
		case <-exit:
			quit = true
			fmt.Printf("exit\n")
			break
		}
		if quit {
			break
		}
	}

	defer func() {
		fmt.Printf("defer exit\n")
		tick.Stop()
		close(exit)
		wg.Done()
	}()
}
