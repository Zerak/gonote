package a

import "gonote/interview/refactor"

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

func (t *A) Add() (ares refactor.Responser, err error) {
	return
}

func (t *A) Sub() (ares refactor.Responser, err error) {
	return
}
