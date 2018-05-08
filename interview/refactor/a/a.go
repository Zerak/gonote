package a

type AResponse struct {
}

func (ar AResponse) GetType() int {
	return 0
}

func (ar AResponse) GetContent() string {
	return "a response"
}

type A struct {
	Content string
}

func (t *A) Add() (ares AResponse, err error) {
	return
}

func (t *A) Sub() (ares AResponse, err error) {
	return
}
