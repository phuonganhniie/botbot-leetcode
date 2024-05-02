package main

import (
	"log"

	"github.com/phuonganhniie/botbot-leetcode/cmd"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())

	// Everyday at 10:00AM
	_, err := c.AddFunc("0 0 10 * * *", func() {
		cmd.Start()
	})
	if err != nil {
		log.Fatalf("Error scheduling cron job: %v", err)
	}

	c.Start() // Start the cron scheduler

	// Use a blocking select statement to keep your application running
	select {}
}
