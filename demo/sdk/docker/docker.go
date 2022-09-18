package docker

import "fmt"

type D struct{}

func (d D) Setup() {
	fmt.Println("must set dns: 8.8.8.8 for docker repo")
}
