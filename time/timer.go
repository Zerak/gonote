package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	//go newTicker()
	go newTimer()

	wg.Wait()
}

func newTicker() {
	fmt.Printf("QueueServer start timer routine\n")
	queueTimer := time.NewTicker(2 * time.Second)
	qTimerCh := queueTimer.C
	for {
		select {
		case <-qTimerCh:
			fmt.Printf("tick\n")
		}
	}

	defer func() {
		fmt.Printf("QueueServer timer routine exit\n")
		queueTimer.Stop()
		wg.Done()
	}()
}

func newTimer() {
	fmt.Printf("QueueServer start timer")
	queueTimer := time.NewTimer(2 * time.Second)
	qTimerCh := queueTimer.C
	for {
		select {
		case <-qTimerCh:
			fmt.Printf("tick\n")
		}
	}
}

func tick() {
	for {
		<-time.After(time.Second * 5)
		fmt.Println(".")
	}
}
