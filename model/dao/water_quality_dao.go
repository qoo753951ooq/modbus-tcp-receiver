package dao

import (
	"modbus-tcp-receiver/db"
	"modbus-tcp-receiver/model"
	"modbus-tcp-receiver/util"
	"strconv"
	"time"
)

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

func (w *WaterQualityLog) UpdateWaterQualityList(eqptListKey string) error {

	key := util.CombineString(eqptListKey, model.Colon, strconv.Itoa(w.Id))
	fieldNames := util.GetStructFieldNames(WaterQualityList{})

	update := make(map[string]interface{}, 0)

	update[fieldNames[0]] = w.Id

	if w.SsStatus != model.Not_Equipment_Field_Status {
		update[fieldNames[1]] = w.SsStatus
	}

	if w.DoStatus != model.Not_Equipment_Field_Status {
		update[fieldNames[2]] = w.DoStatus
	}

	if w.PhStatus != model.Not_Equipment_Field_Status {
		update[fieldNames[3]] = w.PhStatus
	}

	if w.EcStatus != model.Not_Equipment_Field_Status {
		update[fieldNames[4]] = w.EcStatus
	}

	if w.RcStatus != model.Not_Equipment_Field_Status {
		update[fieldNames[5]] = w.RcStatus
	}

	if w.TpStatus != model.Not_Equipment_Field_Status {
		update[fieldNames[6]] = w.TpStatus
	}

	if w.Ss != model.Not_Equipment_Field_Value {
		update[fieldNames[7]] = w.Ss
	}

	if w.Do != model.Not_Equipment_Field_Value {
		update[fieldNames[8]] = w.Do
	}

	if w.Ph != model.Not_Equipment_Field_Value {
		update[fieldNames[9]] = w.Ph
	}

	if w.Ec != model.Not_Equipment_Field_Value {
		update[fieldNames[10]] = w.Ec
	}

	if w.Rc != model.Not_Equipment_Field_Value {
		update[fieldNames[11]] = w.Rc
	}

	if w.Tp != model.Not_Equipment_Field_Value {
		update[fieldNames[12]] = w.Tp
	}

	if w.SsHighAlarm != model.Not_Equipment_Field_Value {
		update[fieldNames[13]] = w.SsHighAlarm
	}

	if w.SsLowAlarm != model.Not_Equipment_Field_Value {
		update[fieldNames[14]] = w.SsLowAlarm
	}

	if w.DoHighAlarm != model.Not_Equipment_Field_Value {
		update[fieldNames[15]] = w.DoHighAlarm
	}

	if w.DoLowAlarm != model.Not_Equipment_Field_Value {
		update[fieldNames[16]] = w.DoLowAlarm
	}

	if w.PhHighAlarm != model.Not_Equipment_Field_Value {
		update[fieldNames[17]] = w.PhHighAlarm
	}

	if w.PhLowAlarm != model.Not_Equipment_Field_Value {
		update[fieldNames[18]] = w.PhLowAlarm
	}

	if w.EcHighAlarm != model.Not_Equipment_Field_Value {
		update[fieldNames[19]] = w.EcHighAlarm
	}

	if w.EcLowAlarm != model.Not_Equipment_Field_Value {
		update[fieldNames[20]] = w.EcLowAlarm
	}

	if w.RcHighAlarm != model.Not_Equipment_Field_Value {
		update[fieldNames[21]] = w.RcHighAlarm
	}

	if w.RcLowAlarm != model.Not_Equipment_Field_Value {
		update[fieldNames[22]] = w.RcLowAlarm
	}

	if w.TpHighAlarm != model.Not_Equipment_Field_Value {
		update[fieldNames[23]] = w.TpHighAlarm
	}

	if w.TpLowAlarm != model.Not_Equipment_Field_Value {
		update[fieldNames[24]] = w.TpLowAlarm
	}

	update[fieldNames[25]] = w.Datatime.Format(util.TimeFormat)

	err := db.RedisHashSet(key, update)

	return err
}
