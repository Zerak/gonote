package main

import (
	"fmt"
	"reflect"
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

	//a := float64(10) /100
	var b int
	b = 10
	a := float64(b) / 100
	fmt.Printf("type a:%v :%v\n", a, reflect.TypeOf(a))

	var servers map[int32]int32
	servers = make(map[int32]int32, 1)
	servers = make(map[int32]int32, 1) // 多次make
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

	if val, ok := servers[1]; ok {
		fmt.Printf("servers have the key of 0 ok:%v val:%v\n", ok, val)
	}

	liveid2Visit := make(map[string][]int)
	liveid2Visit["10010001_liveidRandStr"] = append(liveid2Visit["10010001_liveidRandStr"], 1)
	liveid2Visit["10010001_liveidRandStr"] = append(liveid2Visit["10010001_liveidRandStr"], 2)
	liveid2Visit["10010001_liveidRandStr"] = append(liveid2Visit["10010001_liveidRandStr"], 3)
	liveid2Visit["10010002_liveidRandStr"] = append(liveid2Visit["10010002_liveidRandStr"], 4)
	liveid2Visit["10010001_liveidRandStr"] = append(liveid2Visit["10010001_liveidRandStr"], 1)
	fmt.Println(liveid2Visit)
	for k, v := range liveid2Visit {
		fmt.Printf("k:%v v:%v\n", k, v)
	}

	fmt.Println("kkkk")
	for k := range liveid2Visit {
		fmt.Printf("k:%v\n", k)
	}

	var base int
	var ca float64
	base = 1000000
	ca = 0.12345

	sum := ca * float64(base)
	fmt.Printf("sum:%v intSum:%v\n", sum, int(sum))

}
