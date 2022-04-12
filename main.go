package main

import (
	"modbus-tcp-receiver/conf"
	"modbus-tcp-receiver/service"

	"github.com/robfig/cron/v3"
)

func main() {

	conf := conf.CurrentConfig

	c := cron.New()

	c.AddJob(conf.Param.CronRealTimeSpec, service.SendJob{
		ConfData: conf, SendType: "list"})

	c.Start()

	defer c.Stop()
	select {}
}
