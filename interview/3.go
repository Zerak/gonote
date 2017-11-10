package main

import "fmt"

func fix(s []int) {
	tmp := make([]int,3)
	s = append(s,tmp...)
}

func main() {
	s := []int{1,2,3}
	fix(s)
	fmt.Println(s)
}
