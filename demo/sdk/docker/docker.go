package docker

import "fmt"

type D struct{}

func (d D) Setup() {
	fmt.Println("docker repo must set dns: 8.8.8.8")
}
