package main

import (
	"fmt"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:60000")
	if err != nil {
		fmt.Printf("dialing:%v\n", err)
		return
	}

	var args = "hello rpc"
	var reply string
	err = client.Call("RPC.Ping", args, &reply)
	if err != nil {
		fmt.Printf("RPC.Ping error:%v\n", err)
	} else {
		fmt.Printf("RPC.Ping: %v %v\n", args, reply)
	}

	//err = client.Go("DefaultAuther.Auth", 1, &reply, nil).Done
	//if err != nil {
	//	fmt.Printf("RPC.Auth error:%v\n", err)
	//} else {
	//	fmt.Printf("RPC.Auth: %v %v\n", args, reply)
	//}

	// ping
	tickJob(client)
}

func tickJob(client *rpc.Client) {
	for {
		select {
		case <-time.After(time.Second):
			var args = "."
			var reply string
			err := client.Call("RPC.Ping", args, &reply)
			if err != nil {
				fmt.Printf("ping err:%v\n", err)
			} else {
				fmt.Printf("RPC.Ping:%v\n", reply)
			}
		}
	}
}
