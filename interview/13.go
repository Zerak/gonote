package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func p1() {
	data := []string{"one", "two", "three"}
	for _, v := range data {
		go func() {
			fmt.Println(v)
		}()
	}
}

func p2() {
	data := []string{"one", "two", "three"}
	for _, v := range data {
		vcopy := v //
		go func() {
			fmt.Println(vcopy)
		}()
	}
}
func p3() {
	data := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data {
		go v.print()
	}
}

func main() {
	fmt.Println("p1")
	p1()
	time.Sleep(1 * time.Second)

	fmt.Println("p2")
	p2()
	time.Sleep(1 * time.Second)

	fmt.Println("p3")
	p3()

	time.Sleep(3 * time.Second)
}
