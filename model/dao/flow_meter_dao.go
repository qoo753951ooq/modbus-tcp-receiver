package dao

import "time"

type FlowMeterLog struct {
	Id           int       `orm:"pk;column(id)"`
	FlowStatus   string    `orm:"null;size(50);column(flow_status)"`
	Flow         float64   `orm:"column(flow)"`
	FlowLowAlarm float64   `orm:"null;column(flow_low_alarm)"`
	Datatime     time.Time `orm:"type(datatime)"`
}

type FlowMeterList struct {
	id             int
	flow_status    string
	flow           float64
	flow_low_alarm float64
	datatime       time.Time
}
