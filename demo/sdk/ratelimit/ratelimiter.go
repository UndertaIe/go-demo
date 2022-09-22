package ratelimit

import (
	"fmt"
)

type R struct{}

func (r R) Desc() string {
	return "桶限流算法实现"
}

func (r R) Name() string {
	return "ratelimit8965"
}

func (r R) Repo() string {
	url := "https://github.com/juju/ratelimit"
	fmt.Println(url)
	return url
}
