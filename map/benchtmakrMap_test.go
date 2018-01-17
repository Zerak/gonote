package main

// run with: $ go test --bench=. -test.benchmem .
// run with: $ go test --bench=. -v .
// @see https://twitter.com/karlseguin/status/524452778093977600
import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

const (
	SIZE    = 10
	LOOKUPS = 200
)

// Excluded means that the setup of the lookup variable is created once outside the bench loop
// Included means that the setup of the lookup variable is created new each loop
func BenchmarkStringMapExcluded(b *testing.B) {
	lookup := make(map[string]string, SIZE)
	for i := 0; i < SIZE; i += 1 {
		k := strconv.Itoa(i) // FYI: converts int to string
		lookup[k] = "value"
	}

	rand.Seed(int64(b.N))
	b.ResetTimer()
	get := 0
	for n := 0; n < b.N; n++ {
		//for i := 0; i < LOOKUPS; i++ {
		needle := strconv.Itoa(rand.Intn(SIZE))
		if _, ok := lookup[needle]; ok {
			get++
		}
		//}
	}
	fmt.Printf("StringMap get:%v\n", get)
}
func BenchmarkStringMapIncluded(b *testing.B) {
	rand.Seed(int64(b.N))
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		lookup := make(map[string]string, SIZE)
		for i := 0; i < SIZE; i += 1 {
			k := strconv.Itoa(i) // FYI: converts int to string
			lookup[k] = "value"
		}

		for i := 0; i < LOOKUPS; i++ {
			needle := strconv.Itoa(rand.Intn(SIZE))
			if _, ok := lookup[needle]; ok {
			}
		}
	}
}

func BenchmarkStringSliceExcluded(b *testing.B) {
	lookup := make([]string, SIZE*2)
	for i := 0; i < SIZE; i += 2 {
		lookup[i] = strconv.Itoa(i / 2)
		lookup[i+1] = "value"
	}
	rand.Seed(int64(b.N))
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < LOOKUPS; i++ {
			needle := strconv.Itoa(rand.Intn(SIZE))
			for j := 0; j < SIZE; j += 2 {
				if lookup[j] == needle {

				}
			}
		}
	}
}
func BenchmarkStringSliceIncluded(b *testing.B) {
	rand.Seed(int64(b.N))
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		lookup := make([]string, SIZE*2)
		for i := 0; i < SIZE; i += 2 {
			lookup[i] = strconv.Itoa(i / 2)
			lookup[i+1] = "value"
		}
		for i := 0; i < LOOKUPS; i++ {
			needle := strconv.Itoa(rand.Intn(SIZE))
			for j := 0; j < SIZE; j += 2 {
				if lookup[j] == needle {

				}
			}
		}
	}
}

func BenchmarkStringArraysExcluded(b *testing.B) {
	lookup := [SIZE * 2]string{}
	for i := 0; i < SIZE; i += 2 {
		lookup[i] = strconv.Itoa(i / 2)
		lookup[i+1] = "value"
	}
	rand.Seed(int64(b.N))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < LOOKUPS; i++ {
			needle := strconv.Itoa(rand.Intn(SIZE))
			for j := 0; j < SIZE; j += 2 {
				if lookup[j] == needle {

				}
			}
		}
	}
}
func BenchmarkStringArraysIncluded(b *testing.B) {
	rand.Seed(int64(b.N))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lookup := [SIZE * 2]string{}
		for i := 0; i < SIZE; i += 2 {
			lookup[i] = strconv.Itoa(i / 2)
			lookup[i+1] = "value"
		}
		for i := 0; i < LOOKUPS; i++ {
			needle := strconv.Itoa(rand.Intn(SIZE))
			for j := 0; j < SIZE; j += 2 {
				if lookup[j] == needle {

				}
			}
		}
	}
}

func BenchmarkStructSliceExcluded(b *testing.B) {
	lookup := make([]struct{ key, value string }, SIZE)
	for i := 0; i < SIZE; i++ {
		lookup[i].key = strconv.Itoa(i)
		lookup[i].value = "value"
	}
	rand.Seed(int64(b.N))
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < LOOKUPS; i++ {
			needle := strconv.Itoa(rand.Intn(SIZE))
			for j := 0; j < SIZE; j++ {
				if lookup[j].key == needle {

				}
			}
		}
	}
}
func BenchmarkStructSliceIncluded(b *testing.B) {
	rand.Seed(int64(b.N))
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		lookup := make([]struct{ key, value string }, SIZE)
		for i := 0; i < SIZE; i++ {
			lookup[i].key = strconv.Itoa(i)
			lookup[i].value = "value"
		}
		for i := 0; i < LOOKUPS; i++ {
			needle := strconv.Itoa(rand.Intn(SIZE))
			for j := 0; j < SIZE; j++ {
				if lookup[j].key == needle {

				}
			}
		}
	}
}

func BenchmarkMapInt(b *testing.B) {
	rand.Seed(int64(b.N))
	b.ResetTimer()

	size := 20000000
	look := make(map[int]bool, size)
	for n := 0; n < b.N; n++ {
		key := rand.Intn(size)
		look[key] = true
	}

	total := 0
	for n := 0; n < b.N; n++ {
		key := rand.Intn(size)
		if _, ok := look[key]; ok {
			total++
		}
	}
	fmt.Printf("check:%v\n", total)
}
