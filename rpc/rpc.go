package main

import (
	"fmt"
	"net"
	"net/rpc"
	"strconv"
	"os"
	"os/signal"
	"syscall"
)

const (
	Reply = "reply"
)

type Auther interface {
	Auth(token string) (userId int64, roomId int32)
}

type DefaultAuther struct {
}

func NewDefaultAuther() *DefaultAuther {
	return &DefaultAuther{}
}

func (a *DefaultAuther) Auth(token string) (userId int64, roomId int32) {
	var err error
	if userId, err = strconv.ParseInt(token, 10, 64); err != nil {
		userId = 0
		roomId = -1
	} else {
		roomId = 1 // only for debug
	}
	return
}

// RPC
type RPC struct {
	auther Auther
}

func (r *RPC) Ping(arg string, reply *string) error {
	fmt.Printf("RPC::Ping arg[%v] reply[%v]\n", arg, reply)
	return nil
}

// Connect auth and registe login
func (r *RPC) Connect(arg string, reply *string) (err error) {
	reply = &Reply
	fmt.Printf("RPC::Connect arg[%v] reply[%v]\n", arg, reply)
	r.auther.Auth(arg)
	return
}

// Disconnect notice router offline
func (r *RPC) Disconnect(arg string, reply *string) (err error) {
	fmt.Printf("RPC::Disconnect arg[%v] reply[%v]\n", arg, reply)
	return
}

func rpcListen(network, addr string) {
	l, err := net.Listen(network, addr)
	if err != nil {
		fmt.Printf("net.Listen(\"%s\", \"%s\") error(%v)", network, addr, err)
		panic(err)
	}
	// if process exit, then close the rpc bind
	defer func() {
		fmt.Printf("rpc addr: \"%s\" close", addr)
		if err := l.Close(); err != nil {
			fmt.Printf("listener.Close() error(%v)", err)
		}
	}()
	rpc.Accept(l)
}

func main() {
	c := &RPC{auther: NewDefaultAuther()}
	rpc.Register(c)
	//rpc.HandleHTTP()
	network := "tcp"
	addr := "127.0.0.1:60000"

	go rpcListen(network, addr)

	initSignal()
}

// InitSignal register signals handler.
func initSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	for {
		s := <-c
		fmt.Printf("server get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			return
		case syscall.SIGHUP:
			continue
		default:
			return
		}
	}
}
