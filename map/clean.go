package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	m := make(map[int]string)

	for i := 10; i < 20; i++ {
		m[i] = strconv.Itoa(i + 10)
	}

	for k := range m {
		fmt.Println("k:", k)
	}
	for k, v := range m {
		fmt.Println("k:", k, " val:", v)
	}
	fmt.Println(m)

	s := "abc cde efg abc efg abc cde"
	fmt.Println(counter(s))
}

func counter(s string) map[string]int {
	str := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range str {
		if m[v] == 0 {
			m[v] = 1
		} else {
			m[v] += 1
		}
	}
	return m
}
