package main

import (
	"fmt"
)

func main() {
	n := 0
	count := 0

	word := []string{}
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		str := ""
		n, _ := fmt.Scan(&str)
		if n > 50 {
			continue
		}
		word = append(word,str)
	}
	for _,v := range word{

	}
	fmt.Println(word)
	fmt.Printf("%d\n", count)
}
