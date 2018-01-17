package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	s.Push(2)
	s.Push(3)
	s.Push(2)
	s.Push(1)

	t.Log(s.Empty())
	s.Print()

	t.Log("s min:", s.Min())
	t.Log("s min:", s.Min())
	t.Log("pop:", s.Pop())
	t.Log("s min:", s.Min())
	t.Log("pop:", s.Pop())
	t.Log("empty:",s.Empty())
	t.Log("pop:", s.Pop())
	t.Log("empty:",s.Empty())
}