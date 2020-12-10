package main

import (
	"context"
	"fmt"
	"time"

	"github.com/smallnest/rpcx/client"
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
	c := client.NewClient(client.DefaultOption)
	c.Connect("tcp", "127.0.0.1:8888")
	defer c.Close()

	req := &RpcRequest{Param: "testParam"}
	res := &RpcResponse{}
	//err := c.Call(context.Background(), "RpcRequest", "Request", req, res)
	err := c.Call(context.Background(), "RpcRequest", "Request", req, res)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
