package grpc

import "fmt"

type G struct{}

func (b G) Desc() string {
	return "desc"
}

func (s G) Name() string {
	return "name"
}

func (s G) ReadMe() {
	url := ""
	fmt.Println(url)
}
