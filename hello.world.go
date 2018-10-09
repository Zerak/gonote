package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v8"
)

type MyStruct struct {
	Age int `validate:"gte=3"`
}

func main() {

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
