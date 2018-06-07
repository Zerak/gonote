package main

import (
	"fmt"
	"math/big"
)

func main() {
	m := FormatMoney(0)
	fmt.Println(m)
	return
	fmt.Println(40 / 100000 * 200 * 10000)
	fmt.Println(40 * 200 * 10000 / 100000)
	str := fmt.Sprintf("%.4f\n", float64(1870000)/(10000))
	fmt.Println(str)
	return
	a := big.NewFloat(2.0001)
	b := big.NewFloat(1.00025)
	c := a.Add(a, b)
	d := a.Sub(a, b)
	fmt.Println(c)
	fmt.Println(d)

	aa := big.NewInt(300)
	bb := big.NewInt(200)
	cc := aa.Add(aa, bb)
	dd := aa.Sub(aa, bb)
	fmt.Println(cc)
	fmt.Println(dd)
}

func FormatMoney(num int) string {
	return fmt.Sprintf("%.2f", float64(num)/float64(10000))
}
