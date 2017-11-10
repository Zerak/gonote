package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	s1 := arr[1:4]
	fmt.Println("s1 len:", len(s1), " cap:", cap(s1))
	fmt.Println(s1)

	s2 := s1[2:]
	fmt.Println("s2 len:", len(s2), " cap:", cap(s2))
	fmt.Println(s2)
}
