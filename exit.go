package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Service struct {
	// Other things

	ch        chan bool
	waitGroup *sync.WaitGroup
}

func NewService() *Service {
	s := &Service{
		// Init Other things
		ch:        make(chan bool),
		waitGroup: &sync.WaitGroup{},
	}

	return s
}

func (s *Service) Stop() {
	close(s.ch)
	s.waitGroup.Wait()
}

func (s *Service) Serve() {
	s.waitGroup.Add(1)
	defer s.waitGroup.Done()

	for {
		select {
		case <-s.ch:
			fmt.Println("Serve stopping...")
			return
		default:
		}
		//s.waitGroup.Add(1)
		//go s.anotherServer()
	}
}
func (s *Service) anotherServer() {
	defer s.waitGroup.Done()
	for {
		select {
		case <-s.ch:
			fmt.Println("anotherServer stopping...")
			return
		default:
		}

		// Do something
	}
}

func main() {
	service := NewService()
	go service.Serve()

	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-ch)

	// Stop the service gracefully.
	service.Stop()
}
