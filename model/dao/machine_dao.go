package dao

import (
	"modbus-tcp-receiver/db"
	"modbus-tcp-receiver/model"
	"modbus-tcp-receiver/util"
	"strconv"
	"time"
)

type MachineLog struct {
	Id       int       `orm:"pk;column(id)"`
	Status   string    `orm:"size(50);column(status)"`
	Datatime time.Time `orm:"type(datatime)"`
}

type MachineList struct {
	id       int
	status   string
	datatime time.Time
}

func (m *MachineLog) UpdateMachineList(eqptListKey string) error {

	key := util.CombineString(eqptListKey, model.Colon, strconv.Itoa(m.Id))
	fieldNames := util.GetStructFieldNames(MachineList{})

	update := make(map[string]interface{}, 0)

	update[fieldNames[0]] = m.Id

	if m.Status != model.Not_Equipment_Field_Status {
		update[fieldNames[1]] = m.Status
	}

	update[fieldNames[2]] = m.Datatime.Format(util.TimeFormat)

	err := db.RedisHashSet(key, update)
	return err
}
