package main

import "fmt"

const (
	EPSILON = 1.19209290E-07
)

func main() {
	//var data float64
	var data int64
	data = 0.0000000001

	if data > 0 {
		fmt.Printf("大于")
	} else {
		fmt.Printf("小于")
	}
}
