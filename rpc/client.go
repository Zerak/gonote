package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:60000")
	if err != nil {
		fmt.Printf("dialing:%v\n", err)
	}

	var args = "hello rpc"
	var reply string
	err = client.Call("RPC.Ping", args, &reply)
	if err != nil {
		fmt.Printf("RPC.Ping error:%v\n", err)
	} else {
		fmt.Printf("RPC.Ping: %v %v\n", args, reply)
	}

	err = client.Go("DefaultAuther.Auth", 1, &reply, nil).Done
	if err != nil {
		fmt.Printf("RPC.Auth error:%v\n", err)
	} else {
		fmt.Printf("RPC.Auth: %v %v\n", args, reply)
	}
}
