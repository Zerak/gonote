package main

import (
	"context"
	"fmt"
	"time"

	"github.com/smallnest/rpcx/server"
)

type RpcRequest struct {
	Param string
}

type RpcResponse struct {
	Result string
}

func (r *RpcRequest) Request(ctx context.Context, req *RpcRequest, res *RpcResponse) error {
	res.Result = fmt.Sprintf("%s:%v", req.Param, time.Now().Unix())
	return nil
}

func main() {
	s := server.NewServer()
	defer s.Close()

	s.DisableHTTPGateway = true
	//s.RegisterName("RpcRequest", new(rpcx.RpcRequest), "")
	s.Register(new(RpcRequest), "")
	s.Serve("tcp", "127.0.0.1:8888")
}
