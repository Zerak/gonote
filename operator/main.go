package main

import "fmt"

func main(){
	fmt.Printf("&operator %v\n", 511 & 1)
	fmt.Printf("&operator %v\n", 511 & 2)
	fmt.Printf("&operator %v\n", 511 & 4)
	fmt.Printf("&operator %v\n", 511 & 8)
	fmt.Printf("&operator %v\n", 511 & 16)
	fmt.Printf("&operator %v\n", 511 & 32)
	fmt.Printf("&operator %v\n", 511 & 64)
	fmt.Printf("&operator %v\n", 511 & 128)
	fmt.Printf("&operator %v\n", 511 & 512)
	fmt.Printf("&operator %v\n", 1023 & 1024)
}
