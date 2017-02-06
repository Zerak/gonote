package main

import (
	"fmt"
	"gonote/testAB/a"
	"gonote/testAB/b"
)

func main() {
	A := a.NewA()
	B := b.NewB()
	A.SetB(B)
	B.SetA(A)

	fmt.Println(A.GetN())
	fmt.Println(B.GetN())
}
