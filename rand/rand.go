package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i:= 0; i < 1000; i++  {
		fmt.Println(rand.Int63() ," ", rand.Int())
	}
}
