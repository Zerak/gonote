package main

import "fmt"

// todo
func main() {
	for i := 0; i < 3; i++ {
		func() {
			defer fmt.Println("i:", i)
		}()
	}
}
