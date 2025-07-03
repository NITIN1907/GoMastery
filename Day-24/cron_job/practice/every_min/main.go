package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())

	// Schedule: Every 10 seconds
	c.AddFunc("*/10 * * * * *", func() {
		fmt.Println("Task runs every 10 seconds:", time.Now())
	})

	c.Start()

	select {}
}
