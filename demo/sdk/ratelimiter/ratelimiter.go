package ratelimter

import "fmt"

type R struct{}

func (r R) Desc() string {
	return "sentry示例"
}

func (r R) Name() string {
	return "sentry"
}

func (r R) Repo() string {
	url := "https://github.com/juju/ratelimit"
	fmt.Println(url)
}
