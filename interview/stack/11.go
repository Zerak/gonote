package stack

import "fmt"

// 实现一个特殊的栈，在实现栈的基本功能的基础上，再实现返回栈中最小元素的操作
// Pop、Push、Min操作的时间复杂度都是O（1）
type Stack struct {
	length    int
	stackData []int
	stackMin  []int
}

func (s *Stack) Empty() bool {
	if s.length <= 0 {
		return  true
	}
	return false
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Pop() int {
	if s.length <= 0 {
		panic(fmt.Errorf("empty"))
	}
	res := s.stackData[s.length-1]
	s.stackData = s.stackData[0 : len(s.stackData)-1]
	s.stackMin = s.stackMin[0 : s.length-1]
	s.length--
	return res
}

func (s *Stack) Push(d int) {
	if s.length == 0 {
		s.stackData = append(s.stackData, d)
		s.stackMin = append(s.stackMin, d)
	} else {
		s.stackData = append(s.stackData, d)
		if d <= s.stackMin[0] {
			s.stackMin = append(s.stackMin, d)
		} else {
			s.stackMin = append(s.stackMin, s.Min())
		}
	}
	s.length++
}

// Min get the minest num in Stack
func (s *Stack) Min() int {
	if s.length <= 0 {
		panic(fmt.Errorf("empty"))
	}
	return s.stackMin[s.length-1]
}

func (s *Stack) Print() {
	fmt.Println("\ndata: ")
	for k, v := range s.stackData {
		fmt.Printf("i:%v v:%v\n",k, v)
	}
	fmt.Println("min: ")
	for k, v := range s.stackMin {
		//fmt.Printf("%v ", v)
		fmt.Printf("i:%v v:%v\n",k, v)
	}
}

func NewStack() *Stack {
	return &Stack{}
}