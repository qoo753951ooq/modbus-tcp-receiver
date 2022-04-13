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

func (k *KnifeGateValveLog) InsertKnifeGateValveLog() (int64, error) {

	builder := sqrl.
		Insert(`knife_gate_valve_logs.knife_gate_valve_log`).
		Columns(`id, status, datatime`).
		Values(k.Id, k.Status, k.Datatime)

	if k.Opening != model.Not_Equipment_Field_Value {
		builder.Columns(`opening`)
		builder.Values(k.Opening)
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
