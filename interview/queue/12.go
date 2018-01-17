package queue

import (
	"gonote/interview/stack"
	"fmt"
)

type Queue struct {
	stack.Stack
	stackPush *stack.Stack
	stackPop  *stack.Stack
}

func (q *Queue) Add(d int) {
}

func (q *Queue) Poll() {
}

func (q *Queue) Peek() {
}

func (q *Queue) Print() {
	q.stackPush.Print()
	q.stackPop.Print()
}

func NewQueue() *Queue {
	return &Queue{stackPush: stack.NewStack(), stackPop: stack.NewStack()}
}

func getAndRmLastElement(s *stack.Stack) int {
	res := s.Pop()
	if s.Empty() {
		return res
	} else {
		s.Print()
		//return getAndRmLastElement(s)
		i := getAndRmLastElement(s)
		f := func(a int){
			s.Push(a)
			fmt.Println("push:",a)
		}
		f(i)
		return i
	}
}

func Reverse(s *stack.Stack) {
	if s.Empty() {
		return
	}
	i := getAndRmLastElement(s)
	//Reverse(s)

	f := func(a int){
		s.Push(a)
	}
	f(i)
}

func Fab(n int) int {
	fmt.Println(n)
	if n == 1 {
		return 1
	}
	return n * Fab(n-1)
}