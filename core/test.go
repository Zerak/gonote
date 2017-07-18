package core

import "net"

func main() {
	onStart := func() {

	}
	onEnd := func() {

	}
	handler := func(conn net.Conn) {
		conn
		r := NewReadSession()
		w := NewWriteSession()
		rw := NewRWSession(r, w)

		rw.Run(onStart, onEnd)
	}

	addr := "localhost:6200"
	ListenAndServeTcp(addr, handler, true)
}
