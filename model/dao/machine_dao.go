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

func (m *MachineLog) InsertMachineLog() (int64, error) {

	builder := sqrl.
		Insert(`machine_logs.machine_log`).
		Columns(`id, status, datatime`).
		Values(m.Id, m.Status, m.Datatime)

	sqlStatement, args, _ := builder.ToSql()

	result, err := orm.NewOrm().Raw(sqlStatement, args).Exec()

	if err != nil {
		return 0, err
	} else {
		num, _ := result.RowsAffected()
		return num, nil
	}
}
