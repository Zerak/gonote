package login

import "gonote/cycle/cycle"

type Login struct {
	queue cycle.ServerQueue
	Sum   int
}

func (ls *Login) GetSum() int {
	return ls.Sum
}

func (ls *Login) SetQueuePtr(ptr cycle.ServerQueue) {
	ls.queue = ptr
}
