package dao

import (
	"modbus-tcp-receiver/db"
	"modbus-tcp-receiver/model"
	"modbus-tcp-receiver/util"
	"strconv"
	"time"
)

type FloatLevelSwitchLog struct {
	Id       int       `orm:"pk;column(id)"`
	Level    string    `orm:"size(50);column(level)"`
	Status   string    `orm:"size(50);column(status)"`
	Datatime time.Time `orm:"type(datatime)"`
}

type FloatLevelSwitchList struct {
	id       int
	level    string
	status   string
	datatime time.Time
}

func (f *FloatLevelSwitchLog) UpdateFloatLevelSwitchList(eqptListKey string) error {

	key := util.CombineString(eqptListKey, model.Colon, strconv.Itoa(f.Id))
	fieldNames := util.GetStructFieldNames(FloatLevelSwitchList{})

	update := make(map[string]interface{}, 0)

	update[fieldNames[0]] = f.Id

	if f.Level != model.Not_Equipment_Field_Status {
		update[fieldNames[1]] = f.Level
	}

	if f.Status != model.Not_Equipment_Field_Status {
		update[fieldNames[2]] = f.Status
	}

	update[fieldNames[3]] = f.Datatime.Format(util.TimeFormat)

	err := db.RedisHashSet(key, update)
	return err
}
