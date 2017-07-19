package core

import (
	"fmt"
	"net"
)

var (
	//r  = *ReadSession{}
	w  = *WriteSession{}
	rw = *RWSession{}
)

func main() {
	addr := "localhost:6200"
	ListenAndServeTcp(addr, handleConnection, true)
}

func handleConnection(conn net.Conn) {
	rw = NewRWSession(handleRecvMsg)

	rw.Run(onStart, onEnd)
}

func handleRecvMsg(b []byte) {
	fmt.Println("recv str:", string(b))
	fmt.Println("recv byte:", b)

	rw.Send(BytesPacket{"str"})
}

func onStart() {
	fmt.Println("server started")
}

func onEnd() {
	fmt.Println("server ended")
}
