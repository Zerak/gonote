package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v8"
)

type MyStruct struct {
	Age int `validate:"gte=3"`
}

func main() {
	maxX := 1000
	// maxY := 1000
	pixelsPerRow := 50
	getRow := func(row, col int) int {
		x := row / maxX
		y := (maxX / pixelsPerRow) * col
		return x + y
	}
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			idx := x + y + 1
			fix := "\n"
			if y >= 10 && y%10 == 0 {
				fix = "\n"
			}
			fmt.Printf("row_idx:%v idx:%v row:%v col:%v %v", idx, getRow(x, y), x, y, fix)
		}
	}
	return

	config := &validator.Config{TagName: "validate"}
	valid := validator.New(config)

	s := MyStruct{Age: 3}
	err := valid.Struct(s)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("ok")
	}

}
