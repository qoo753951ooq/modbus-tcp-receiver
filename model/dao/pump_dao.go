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

type PumpLog struct {
	Id       int       `orm:"pk;column(id)"`
	Status   string    `orm:"size(50);column(status)"`
	Datatime time.Time `orm:"type(datatime)"`
}

type PumpList struct {
	id       int
	status   string
	datatime time.Time
}

func (l *PumpLog) UpdatePumpList(eqptListKey string) error {

	key := util.CombineString(eqptListKey, model.Colon, strconv.Itoa(l.Id))
	fieldNames := util.GetStructFieldNames(PumpList{})

	update := make(map[string]interface{}, 0)

	update[fieldNames[0]] = l.Id

	if l.Status != model.Not_Equipment_Field_Status {
		update[fieldNames[1]] = l.Status
	}

	update[fieldNames[2]] = l.Datatime.Format(util.TimeFormat)

	err := db.RedisHashSet(key, update)
	return err
}

func (l *PumpLog) InsertPumpLog() (int64, error) {

	builder := sqrl.
		Insert(`pump_logs.pump_log`).
		Columns(`id, status, datatime`).
		Values(l.Id, l.Status, l.Datatime)

	sqlStatement, args, _ := builder.ToSql()

	result, err := orm.NewOrm().Raw(sqlStatement, args).Exec()

	if err != nil {
		return 0, err
	} else {
		num, _ := result.RowsAffected()
		return num, nil
	}
}
