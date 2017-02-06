package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	argError = errors.New("参数错误")
)

func main() {
	str := ",1,2,3,3,4,"
	str = strings.TrimRight(str, ",")
	fmt.Println("trim->", str)
	str = strings.TrimLeft(str, ",")
	fmt.Println("trim->", str)
	fmt.Printf("arg:%+v\n", argError)
	for i := 10000000; i < 10000099; i++ {
		fmt.Println(encodeUserId(i))
	}
}

func encodeUserId(uid int) string {
	base := 10000000
	if uid < base {
		return "100***00"
	}
	tail := uid % base
	return fmt.Sprintf("100***%02d %v", tail, tail)
}
