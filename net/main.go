package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("192.168.1.1")
	ipv4 := ip.To4()
	if ipv4 == nil {
		// not ipv4
		fmt.Println("not ipv4")
	} else {
		fmt.Println("ipv4")
	}
}
