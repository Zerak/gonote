package main

import "fmt"

func main(){
	fmt.Printf("&operator %v\n", 447 & 1)
	fmt.Printf("&operator %v\n", 447 & 2)
	fmt.Printf("&operator %v\n", 447 & 4)
	fmt.Printf("&operator %v\n", 447 & 8)
	fmt.Printf("&operator %v\n", 447 & 16)
	fmt.Printf("&operator %v\n", 447 & 32)
	fmt.Printf("&operator %v\n", 447 & 64)
	fmt.Printf("&operator %v\n", 447 & 128)
	fmt.Printf("&operator %v\n", 447 & 256)
	fmt.Printf("&operator %v\n", 447 & 512)
	fmt.Printf("&operator %v\n", 1023 & 1024)
}
