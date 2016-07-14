package main

import (
	"container/list"
	"fmt"
)

var sMap map[int]bool

func main() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(21)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(2)
	l.PushBack(7)
	l.PushBack(8)
	l.PushBack(9)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("befor")

	var n *list.Element
	update := false
	for e := l.Front(); e != nil; e = n {
		n = e.Next()
		if update {
			fmt.Println(e.Value)
		}
		if e.Value == 2 {
			l.Remove(e)
			update = true
		}
		fmt.Printf("... [%v]\n", update)
	}

	fmt.Println("remove")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
