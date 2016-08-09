package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	var once sync.Once
	for i := 0; i < 10; i++ {
		fmt.Printf("new goroutine i[%v]...\n", i)
		once.Do(func() {
			wg.Add(1)
			go serverQueue(i)
		})
		fmt.Printf("new goroutine\n")
	}
	wg.Wait()
	fmt.Printf("main exit\n")
}

func serverQueue(id int) {
	queueTimer := time.NewTicker(2 * time.Second)
	qTimerCh := queueTimer.C
	for {
		select {
		case <-qTimerCh:
			fmt.Printf("tick id[%v]\n", id)
		default:
			fmt.Printf(".")
		}
	}

	defer func() {
		wg.Done()
	}()
}
