package main

import (
	"log"

	"github.com/phuonganhniie/botbot-leetcode/cmd"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	_, err := c.AddFunc("0 10 * * *", cmd.Start) // Everyday at 10:00AM
	if err != nil {
		log.Fatalf("Error scheduling cron job: %v", err)
	}

	c.Start()

	select {}
}
