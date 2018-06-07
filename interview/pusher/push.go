package pusher

import (
	"sync"
)

type Responser interface {
	GetType() int
	GetContent() string
}

type Pusher interface {
	PushToAll() (Responser, error)
	PushToUser() (Responser, error)
}

var (
	pusherMu sync.RWMutex
	pushers  = make(map[string]Pusher)
)

func register(name string, push Pusher) {
	pusherMu.Lock()
	defer pusherMu.Unlock()
	if push == nil {
		panic("pusher: Register push is nil")
	}
	if _, dup := pushers[name]; dup {
		panic("pusher: Register called twice for push" + name)
	}
	pushers[name] = push
}

type Option struct {
	Type int
	Name string
}

func NewPusher(o *Option) Pusher {
	switch o.Type {
	case 1:
		register(o.Name, Ali)
	case 2:
		register(o.Name, Ali)
	}
}
