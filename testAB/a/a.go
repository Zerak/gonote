package a

import (
	"gonote/testAB/er"
)

var n = 2

func (a *A) GetN() string {
	return a.b.GetN()
}

/***************************************/
type A struct {
	b er.B
}

func (a *A) SetB(b er.B) {
	a.b = b
}

func NewA() *A {
	return new(A)
}
