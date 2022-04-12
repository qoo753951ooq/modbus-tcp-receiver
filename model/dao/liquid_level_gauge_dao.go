package dao

import (
	"modbus-tcp-receiver/db"
	"modbus-tcp-receiver/model"
	"modbus-tcp-receiver/util"
	"strconv"
	"time"
)

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

func (l *LiquidLevelGaugeLog) UpdateLiquidLevelGaugeList(eqptListKey string) error {

	key := util.CombineString(eqptListKey, model.Colon, strconv.Itoa(l.Id))
	fieldNames := util.GetStructFieldNames(LiquidLevelGaugeList{})

	update := make(map[string]interface{}, 0)

	update[fieldNames[0]] = l.Id

	if l.Level != model.Not_Equipment_Field_Status {
		update[fieldNames[1]] = l.Level
	}

	if l.Status != model.Not_Equipment_Field_Status {
		update[fieldNames[2]] = l.Status
	}

	if l.Liquid != model.Not_Equipment_Field_Value {
		update[fieldNames[3]] = l.Liquid
	}

	if l.LiquidHighAlarm != model.Not_Equipment_Field_Value {
		update[fieldNames[4]] = l.LiquidHighAlarm
	}

	if l.LiquidLowAlarm != model.Not_Equipment_Field_Value {
		update[fieldNames[5]] = l.LiquidLowAlarm
	}

	update[fieldNames[6]] = l.Datatime.Format(util.TimeFormat)

	err := db.RedisHashSet(key, update)
	return err
}
