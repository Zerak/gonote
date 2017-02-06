package main

import "fmt"

func main() {
	//cas := true
	cas := false
	if cas {
		return
	}
	defer func() {
		fmt.Printf("defer")
	}()
}
