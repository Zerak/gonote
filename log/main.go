package main

import (
	"fmt"
	"sync"
)

func main() {
	m := new(sync.Map)
	for i := 0; i < 100; i++ {
		m.Store(fmt.Sprintf("%v", i), i)
	}

	m.Range(func(key, value interface{}) bool {
		fmt.Println("Key:", key)
		fmt.Println("Val:", value)
		return true
	})
}
