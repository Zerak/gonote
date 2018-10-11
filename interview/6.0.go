package main

import "fmt"

// todo
func main() {
	for i := 0; i < 3; i++ {
		defer fmt.Println("i:", i)
	}
}
