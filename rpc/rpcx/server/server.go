package main

import (
	"github.com/smallnest/rpcx/server"
)

func main() {
	s := server.NewServer()
	defer s.Close()

	s.DisableHTTPGateway = true
	s.RegisterName("RpcRequest", new(RpcRequest), "")
	s.Serve("tcp", "127.0.0.1:8888")
}
