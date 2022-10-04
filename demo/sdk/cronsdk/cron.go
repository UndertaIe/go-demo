package cronsdk

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func (cc C) DemoOfCron() {
	c := cron.New()
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()
	c.AddFunc("@daily", func() { fmt.Println("Every day") })
	c.Stop() // Stop the scheduler (does not stop any jobs already running).
}
