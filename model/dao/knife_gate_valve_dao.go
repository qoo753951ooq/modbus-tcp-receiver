package dao

import (
	"modbus-tcp-receiver/db"
	"modbus-tcp-receiver/model"
	"modbus-tcp-receiver/util"
	"strconv"
	"time"
)

type KnifeGateValveLog struct {
	Id       int       `orm:"pk;column(id)"`
	Status   string    `orm:"size(50);column(status)"`
	Opening  float64   `orm:"null;column(opening)"`
	Datatime time.Time `orm:"type(datatime)"`
}

type KnifeGateValveList struct {
	id       int
	status   string
	opening  float64
	datatime time.Time
}

func (k *KnifeGateValveLog) UpdateKnifeGateValveList(eqptListKey string) error {

	key := util.CombineString(eqptListKey, model.Colon, strconv.Itoa(k.Id))
	fieldNames := util.GetStructFieldNames(KnifeGateValveList{})

	update := make(map[string]interface{}, 0)

	update[fieldNames[0]] = k.Id

	if k.Status != model.Not_Equipment_Field_Status {
		update[fieldNames[1]] = k.Status
	}

	if k.Opening != model.Not_Equipment_Field_Value {
		update[fieldNames[2]] = k.Opening
	}

	update[fieldNames[3]] = k.Datatime.Format(util.TimeFormat)

	err := db.RedisHashSet(key, update)
	return err
}
