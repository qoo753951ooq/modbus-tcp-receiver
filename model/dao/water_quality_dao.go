package dao

import "time"

type WaterQualityLog struct {
	Id          int       `orm:"pk;column(id)"`
	SsStatus    string    `orm:"null;size(50);column(ss_status)"`
	DoStatus    string    `orm:"null;size(50);column(do_status)"`
	PhStatus    string    `orm:"null;size(50);column(ph_status)"`
	EcStatus    string    `orm:"null;size(50);column(ec_status)"`
	RcStatus    string    `orm:"null;size(50);column(rc_status)"`
	TpStatus    string    `orm:"null;size(50);column(tp_status)"`
	Ss          float64   `orm:"null;column(ss)"`
	Do          float64   `orm:"null;column(do)"`
	Ph          float64   `orm:"null;column(ph)"`
	Ec          float64   `orm:"null;column(ec)"`
	Rc          float64   `orm:"null;column(rc)"`
	Tp          float64   `orm:"null;column(tp)"`
	SsHighAlarm float64   `orm:"null;column(ss_high_alarm)"`
	SsLowAlarm  float64   `orm:"null;column(ss_low_alarm)"`
	DoHighAlarm float64   `orm:"null;column(do_high_alarm)"`
	DoLowAlarm  float64   `orm:"null;column(do_low_alarm)"`
	PhHighAlarm float64   `orm:"null;column(ph_high_alarm)"`
	PhLowAlarm  float64   `orm:"null;column(ph_low_alarm)"`
	EcHighAlarm float64   `orm:"null;column(ec_high_alarm)"`
	EcLowAlarm  float64   `orm:"null;column(ec_low_alarm)"`
	RcHighAlarm float64   `orm:"null;column(rc_high_alarm)"`
	RcLowAlarm  float64   `orm:"null;column(rc_low_alarm)"`
	TpHighAlarm float64   `orm:"null;column(tp_high_alarm)"`
	TpLowAlarm  float64   `orm:"null;column(tp_low_alarm)"`
	Datatime    time.Time `orm:"type(datatime)"`
}

type WaterQualityList struct {
	id            int
	ss_status     string
	do_status     string
	ph_status     string
	ec_status     string
	rc_status     string
	tp_status     string
	ss            float64
	do            float64
	ph            float64
	ec            float64
	rc            float64
	tp            float64
	ss_high_alarm float64
	ss_low_alarm  float64
	do_high_alarm float64
	do_low_alarm  float64
	ph_high_alarm float64
	ph_low_alarm  float64
	ec_high_alarm float64
	ec_low_alarm  float64
	rc_high_alarm float64
	rc_low_alarm  float64
	tp_high_alarm float64
	tp_low_alarm  float64
	datatime      time.Time
}

type WaterQualityStatus struct {
	Ss_status string
	Do_status string
	Ph_status string
	Ec_status string
	Rc_status string
	Tp_status string
}
