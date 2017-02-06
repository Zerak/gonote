package main

import "fmt"

func main() {
	defer fmt.Println("a")
	defer fmt.Println("b")
	if true {
		fmt.Println("c1")
		defer fmt.Println("c") // 后入栈，先执行
		fmt.Println("c2")
	}
	fmt.Println("d")
}
