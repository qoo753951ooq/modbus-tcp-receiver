package main

import (
	"modbus-tcp-receiver/conf"
	"modbus-tcp-receiver/db"
	"modbus-tcp-receiver/service"

	"github.com/robfig/cron/v3"
)

func main() {

	conf := conf.CurrentConfig
	db.InitRedisDB(conf.Param.Redis.Ip, conf.Param.Redis.Pwd, conf.Param.Redis.Db)
	db.InitPostgresDB(conf.Param.Postgres.DataSource)

	service.CheckRedisEqptKey(conf)

	c := cron.New()

	c.AddJob(conf.Param.CronRealTimeSpec, service.SendJob{
		ConfData: conf, SendType: "list"})

	c.AddJob(conf.Param.CronHistorySpec, service.SendJob{
		ConfData: conf, SendType: "log"})

	c.Start()

	defer c.Stop()
	select {}
}
