package main

import (
	"fmt"
	"strings"
)

func main() {
	method := "/passport/getuserinfo 是/"
	s := strings.Trim(method, "/")
	fmt.Println(s)

	st := strings.Title(s)
	fmt.Println(st)
}
