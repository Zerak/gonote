package main

import (
	"fmt"
	"gonote/cycle/login"
	"gonote/cycle/queue"
)

func main() {
	li := new(login.Login)
	li.Sum = 10

	qu := new(queue.Queue)
	qu.Sum = 2

	//li := login.Login{Sum:10}
	//qu := queue.Queue{Sum:2}

	li.SetQueuePtr(qu)
	qu.SetLoginPtr(li)

	fmt.Printf("qu getsum[%v]\n", qu.GetSum())
	fmt.Printf("li getsum[%v]\n", li.GetSum())
}
