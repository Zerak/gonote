package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var tem = `
package handler

import "fmt"

// NewLoginHandler login handler
func NewLoginHandler(taskID int) (Handler, error) {
	return &Login{TaskID: taskID}, nil
}

type Login struct {
	TaskID int
}

func (l *Login) Process(e Event) error {
	fmt.Println("login:", e)
	return nil
}
`

func main() {
	fname := flag.String("f", "Handler", "your handler name")
	dir := flag.String("d", "./", "your file where to place")
	flag.Parse()

	f, err := os.Open(*fname)
	if err != nil {
		panic(err)
	}

	tem = fmt.Sprintf("")
	n, err := f.WriteString(tem)
	if err != nil {
		panic(err)
	}
	fmt.Printf("write byte:%v", n)

	str := fmt.Sprintf("cp %v %v", fname, dir)
	err = exec.Command(str).Run()
	if err != nil {
		panic(err)
	}
}
