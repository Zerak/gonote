package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Map map[string]string

var m atomic.Value
var mu sync.Mutex

func main() {
	read := func(key string) (val string) {
		if v, ok := m.Load().(Map); ok {
			return v[key]
		}
		return ""
	}

	insert := func(key, val string) {
		mu.Lock()
		defer mu.Unlock()
		oldMap := m.Load().(Map)
		newMap := make(Map)
		for k, v := range oldMap {
			newMap[k] = v
		}
		newMap[key] = val
		m.Store(newMap)
	}

	k := "key"
	v := read(k)
	fmt.Println(v)

	key := "key"
	val := "val"
	insert(key, val)
}
