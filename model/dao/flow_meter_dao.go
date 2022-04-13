package dao

import (
	"modbus-tcp-receiver/db"
	"modbus-tcp-receiver/model"
	"modbus-tcp-receiver/util"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/elgris/sqrl"
)

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

func (f *FlowMeterLog) UpdateFlowMeterList(eqptListKey string) error {

	key := util.CombineString(eqptListKey, model.Colon, strconv.Itoa(f.Id))
	fieldNames := util.GetStructFieldNames(FlowMeterList{})

	update := make(map[string]interface{}, 0)

	update[fieldNames[0]] = f.Id

	if f.FlowStatus != model.Not_Equipment_Field_Status {
		update[fieldNames[1]] = f.FlowStatus
	}

	if f.Flow != model.Not_Equipment_Field_Value {
		update[fieldNames[2]] = f.Flow
	}

	if f.FlowLowAlarm != model.Not_Equipment_Field_Value {
		update[fieldNames[3]] = f.FlowLowAlarm
	}

	update[fieldNames[4]] = f.Datatime.Format(util.TimeFormat)

	err := db.RedisHashSet(key, update)

	return err
}

func (f *FlowMeterLog) InsertFlowMeterLog() (int64, error) {

	builder := sqrl.
		Insert(`flow_meter_logs.flow_meter_log`).
		Columns(`id, flow, datatime`).
		Values(f.Id, f.Flow, f.Datatime)

	if f.FlowStatus != model.Not_Equipment_Field_Status {
		builder.Columns(`flow_status`)
		builder.Values(f.FlowStatus)
	}

	if f.FlowLowAlarm != model.Not_Equipment_Field_Value {
		builder.Columns(`flow_low_alarm`)
		builder.Values(f.FlowLowAlarm)
	}

	sqlStatement, args, _ := builder.ToSql()

	result, err := orm.NewOrm().Raw(sqlStatement, args).Exec()

	if err != nil {
		return 0, err
	} else {
		num, _ := result.RowsAffected()
		return num, nil
	}
}
