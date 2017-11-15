package main

import "fmt"

func main() {
	a := 1
	b := 2
	defer fmt.Println("a:", a, " b:", b)
	a = 3
	b = 4
}
