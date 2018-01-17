package main

import "fmt"

func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	s1 := arr[0:4] // startIndex,endIndex-1	len=endIndex-startIndex cap=len(arr) - startIndex
	fmt.Println("s1 len:", len(s1), " cap:", cap(s1))
	fmt.Println(s1)

	s2 := s1[2:]
	fmt.Println("s2 len:", len(s2), " cap:", cap(s2))
	fmt.Println(s2)
}
