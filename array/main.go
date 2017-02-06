package main

import "fmt"

func main() {
	robots := make([][]int, 50)

	for i := 0; i < 50; i++ {
		robots[i][i] = i
		robots[i][i] = i * 20
	}
	fmt.Println(robots)
}
