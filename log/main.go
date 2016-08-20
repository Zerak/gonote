package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	glog.Info("log")
}
