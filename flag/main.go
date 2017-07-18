package main

import (
	"flag"
	"fmt"
)

type Config struct {
	Name string
	Val  string
}

func (c Config) String() string {
	return fmt.Sprintf("name:%s val:%s", c.Name, c.Val)
}

// go run main.go -name name -val val
func main() {
	conf := Config{}
	flag.CommandLine.StringVar(&conf.Name, "name", "the name", "usage of cmdline")
	flag.CommandLine.StringVar(&conf.Val, "val", "the val", "usage of cmdline")
	flag.Parse()

	fmt.Println(conf)
}
