package main

import "fmt"

func main() {
	ch := make(chan int, 100)

	for i := 0; i < 10; i++ {
		ch <- 1
	}

	fmt.Println(len(ch))
	fmt.Println(cap(ch))
}
