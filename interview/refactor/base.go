package refactor

import (
	"gonote/interview/refactor/a"
	"gonote/interview/refactor/b"
)

type Responser interface {
	GetType() int
	GetContent() string
}

type Baser interface {
	Add() (Responser, error)
	Sub() (Responser, error)
}

func getClass(tpe int) Baser {
	switch tpe {
	case 1:
		return &a.A{Content: "a"}
	case 2:
		return &b.B{Content: "b"}
	}
	return nil
}

func main() {
	getClass(1)
}
