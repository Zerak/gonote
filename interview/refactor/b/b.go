package b

type BResponse struct {
}

func (br BResponse) GetType() int {
	return 0
}

func (br BResponse) GetContent() string {
	return "b response"
}

type B struct {
	Content string
}

func (t *B) Add() (ares BResponse, err error) {
	return
}

func (t *B) Sub() (ares BResponse, err error) {
	return
}
