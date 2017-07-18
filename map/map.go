package main

import "fmt"

func main() {
	type Key [64]byte
	type Value struct {
		Name      [32]byte
		Balance   uint64
		Timestamp int64
	}
	m := make(map[Key]Value, 1e8)

	fmt.Println("len:", len(m), " m:", m)
}
