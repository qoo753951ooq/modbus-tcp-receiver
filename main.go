package main

import (
	"github.com/robfig/cron/v3"
)

func main() {

	c := cron.New()

	c.Start()

	defer c.Stop()
	select {}
}
