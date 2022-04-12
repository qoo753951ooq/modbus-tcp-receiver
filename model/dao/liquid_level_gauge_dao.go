package dao

import "time"

type LiquidLevelGaugeLog struct {
	Id              int       `orm:"pk;column(id)"`
	Level           string    `orm:"size(50);column(level)"`
	Status          string    `orm:"size(50);column(status)"`
	Liquid          float64   `orm:"column(liquid)"`
	LiquidHighAlarm float64   `orm:"null;column(liquid_high_alarm)"`
	LiquidLowAlarm  float64   `orm:"null;column(liquid_low_alarm)"`
	Datatime        time.Time `orm:"type(datatime)"`
}

type LiquidLevelGaugeList struct {
	id                int
	level             string
	status            string
	liquid            float64
	liquid_high_alarm float64
	liquid_low_alarm  float64
	datatime          time.Time
}
