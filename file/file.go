package main

import (
	"fmt"
	"path/filepath"
	"os"
)

func main() {
	dir, file := filepath.Split("./file.go")
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir,os.FileMode(0775)); err != nil{
			fmt.Println("mkdir err")
			return
		}
		fmt.Printf("dir:%v file:%v\n", dir, file)
	} else {
		fmt.Println("dir empty")
	}
}
