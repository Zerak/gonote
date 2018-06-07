package main

import (
	"fmt"
	"math/big"
)

const (
	EPSILON = 1.19209290E-07
)

func main() {
	a := big.NewFloat(2.01234567890123456789)
	b := big.NewFloat(1.09876543210987654321)

	c := a.Add(a, b)
	fmt.Printf("%v\n", c.String())

	////var data float64
	//var data int64
	//data = 0.0000000001
	//
	//if data > 0 {
	//	fmt.Printf("大于")
	//} else {
	//	fmt.Printf("小于")
	//}

	/*
		var a = 81
		b := float32(a) / 10
		fmt.Println("b:", b, reflect.TypeOf(b))
		var aa []int
		//aa = make([]int,0)

		for k, v := range aa {
			fmt.Printf("k:%v v:%v\n", k, v)
		}

		like := 10
		unlike := 3
		result := fmt.Sprintf("%.2v%%", float32(like)/float32(like+unlike)*100)
		fmt.Println(result)
	*/
}
