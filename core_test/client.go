package main

import (
	"fmt"
	"net"

	"gonote/core"
	"os/signal"
)

func main() {
	addr := "localhost:6200"
	err := core.ListenAndServeTcp(addr, handleConnection, true)
	if err != nil {
		fmt.Println("listen and serve error:", err)
	}

	signal.Ignore()
	signal.Notify()
}

func handleConnection(conn net.Conn) {
	fmt.Println("connected")
	rw := core.NewRWSession(handleRecvMsg)
	rw.Run(onStart, onEnd)
}

func handleRecvMsg(b []byte) {
	fmt.Println("recv str:", string(b))
	fmt.Println("recv byte:", b)

	//rw.Send(core.BytesPacket{byte("str")})
}

func onStart() {
	fmt.Println("server started")
}

func onEnd() {
	fmt.Println("server ended")
}
