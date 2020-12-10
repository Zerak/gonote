package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:10203")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	rend := make(chan struct{})
	wend := make(chan struct{})

	go writeLoop(conn, wend)
	go startLoop(conn, rend, wend)

	<-rend
	<-wend
	fmt.Printf("end")
}

func startLoop(conn *net.TCPConn, rend, wend chan<- struct{}) {
	// first login
	login := []byte("@login 5 55e1ff62b2c8463d101230rKNljIf273\n")
	n, err := conn.Write(login)
	if err != nil {
		rend <- struct{}{}
		return
	}
	fmt.Printf("login write:%v src:%v\n", n, len(login))

	// login res
	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Printf("login res err:%v\n", err)
		return
	}
	fmt.Printf("login res:%v\n", string(buf))

	for {
		select {
		case <-time.After(time.Second):
			ping := []byte("@ping")
			n, err := conn.Write(ping)
			if err != nil {
				panic(err)
			}
			fmt.Println("n:", n)

			buf, err := ioutil.ReadAll(conn)
			if err != nil {
				fmt.Printf("ping res err:%v\n", err)
				return
			}
			fmt.Printf("ping res:%v\n", string(buf))

		}
	}
	rend <- struct{}{}
}

func writeLoop(conn *net.TCPConn, wend chan<- struct{}) {
	wend <- struct{}{}
}
