package main

import (
	"fmt"
)

func main() {
	//mp := map[string]int{"a":1,"b":2,"c":3}
	//
	//val, ok := mp["d"]
	//fmt.Printf("ok[%v], val[%v]]\n", ok, val)
	//
	//for key, val := range mp{
	//	fmt.Printf("key[%v] val[%v]\n", key, val)
	//}
	//
	//var pp *int
	//if pp != nil {
	//	fmt.Printf("pp not nill\n")
	//}else {
	//	fmt.Printf("pp nill\n")
	//}

	var servers map[int32]int32
	servers = make(map[int32]int32, 1)
	for k, v := range servers {
		fmt.Printf("key:%v val:%v \n", k, v)
	}
	servers[0] = 10
	servers[1] = 11
	servers[2] = 12
	fmt.Printf("server:%v len:%v\n", servers, len(servers))

	var arr []int32
	arr = make([]int32, 3, 5)
	for i, v := range arr {
		fmt.Printf("i:%v val:%v\n", i, v)
	}
	fmt.Printf("arr:%v len:%v cap:%v\n", arr, len(arr), cap(arr))
}
