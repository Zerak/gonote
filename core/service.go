package core

import (
	"crypto/tls"
	"net"
)

func listenTCP(addrStr string) (net.Listener, error) {
	addr, err := net.ResolveTCPAddr("tcp4", addrStr)
	if err != nil {
		return nil, err
	}

	listener, err := net.ListenTCP("tcp", addr)
	return listener, err
}

func serve(listener net.Listener, handler func(net.Conn), async bool) error {
	serveFunc := func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			go handler(conn)
		}
	}

	if async {
		go serveFunc()
	} else {
		serveFunc()
	}

	return nil
}

func ListenAndServeTcp(addrStr string, handler func(net.Conn), async bool, certs ...tls.Certificate) error {
	var (
		listener net.Listener
		err      error
	)

	if len(certs) > 0 {
		config := &tls.Config{Certificates: certs}
		listener, err = tls.Listen("tcp", addrStr, config)
	} else {
		listener, err = listenTCP(addrStr)
	}

	if err == nil {
		serve(listener, handler, async)
	}

	return err
}
