package b

import (
	"gonote/testAB/er"
)

var n = 3

func (b *B) GetN() string {
	return "this is packet b func GetN\n"
}

/**************************************/
type B struct {
	a er.A
}

func NewB() *B {
	return new(B)
}
func (b *B) SetA(a er.A) {
	b.a = a
}
