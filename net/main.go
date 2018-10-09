package main

import (
	"fmt"
	"net"
	"net/http"
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

	http.Handle("index", NewHttp())
	http.ListenAndServe("localhost", NewHttp())
}

type DefaultHttpServer struct {
}

func (h *DefaultHttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func NewHttp() *DefaultHttpServer {
	return &DefaultHttpServer{}
}
