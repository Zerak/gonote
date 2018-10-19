package main

import (
	"fmt"
	"math/rand"
)

func main() {
	mmmm := make(map[int]int64)
	mmmm[123] = 12340
	for k, v := range mmmm {
		fmt.Println(k, v)
	}
	return

	m2 := make(map[string]string)
	m2["2"] = "21"
	if v, ok := m2["2"]; ok {
		fmt.Println(v)
	}

	//type Key [64]byte
	//type Value struct {
	//	Name      [32]byte
	//	Balance   uint64
	//	Timestamp int64
	//}
	//m := make(map[Key]Value, 1e8)
	//
	//fmt.Println("len:", len(m), " m:", m)

	size := 200
	tm := make(map[int]bool, size)
	for i := 0; i < size; i++ {
		key := rand.Intn(size)
		//fmt.Printf("key:%v\n", key)
		if _, ok := tm[key]; ok {
			//t2 := tm[key]
			//fmt.Printf("key:%v ok:%v t2:%v\n", key, ok, t2)
			continue
		}
		tm[key] = true
	}

	fmt.Printf("begin check...\n")
	total := 0
	for i := 0; i < size; i++ {
		key := rand.Intn(size)
		if _, ok := tm[key]; ok {
			//fmt.Printf("get key:%v\n", key)
			total++
		}
	}
	fmt.Printf("tm len:%v check total:%v\n", len(tm), total)

	fmt.Printf("pseudorand test...\n")
	for i := 0; i < 10; i++ {
		rand.Seed(int64(i))
		fmt.Println(rand.Intn(10))
	}

	m := make(map[string]int)
	m["101:10001"] = 100

	v, _ := m["101:10001"]
	fmt.Println("value:", v, " ok:", "")
}
