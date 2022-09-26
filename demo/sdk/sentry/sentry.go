package sentry

import (
	"errors"
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
)

type S struct{}

func (s S) Desc() string {
	return "sentry示例"
}

func (s S) Name() string {
	return "sentry"
}

func (s S) Setup() {
	url := "https://hub.docker.com/_/sentry#:~:text=How%20to%20setup%20a%20full%20Sentry%20instance"
	fmt.Println(url)
}

// visit http://192.168.3.200:9000/sentry/
func (s S) DemoOfSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   "http://23dbc042add54e3db870c63fa03d470e@192.168.3.200:9000/2",
		Debug: true,
	})
	fmt.Println(err)
	sentry.CaptureException(errors.New("my error"))
	sentry.Flush(time.Second * 5)
}
