package queue

import (
	"testing"
	"gonote/interview/stack"
	"fmt"
)

func TestQueue(t *testing.T) {
		q := NewQueue()

		q.Add(1)
		q.Add(1)
		q.Add(1)
		q.Add(1)
		q.Add(1)

		q.Poll()
		q.Peek()

		q.Print()
}

func TestReverse(t *testing.T) {
	s := stack.NewStack()
	s.Push(3)
	s.Push(2)
	s.Push(1)

	fmt.Println("before reverse:")
	s.Print()

	fmt.Println("reverse...")
	Reverse(s)
	fmt.Println("after reverse:")
	s.Print()
}

func TestFab(t *testing.T) {
	t.Log(Fab(3))
}
