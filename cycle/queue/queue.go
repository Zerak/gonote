package queue

import "gonote/cycle/cycle"

type Queue struct {
	login cycle.ServerLogin
	Sum   int
}

func (qs *Queue) GetSum() int {
	return qs.Sum
}

func (qs *Queue) SetLoginPtr(ptr cycle.ServerLogin) {
	qs.login = ptr
}
