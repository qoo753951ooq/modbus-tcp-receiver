package service

import (
	"modbus-tcp-receiver/conf"
	"modbus-tcp-receiver/conf/md"
	"modbus-tcp-receiver/db"
	"modbus-tcp-receiver/model"
	"modbus-tcp-receiver/model/dao"
	"modbus-tcp-receiver/util"
	"reflect"
	"strconv"
)

func CheckRedisEqptKey(confData conf.Configuration) {

	eqptKeys := confData.Param.Redis.EquipmentKeys

	newEquipmentKey(eqptKeys[0], confData.KnifeGateValves, dao.KnifeGateValveLog{})
	newEquipmentKey(eqptKeys[1], confData.Pumps, dao.PumpLog{})
	newEquipmentKey(eqptKeys[2], confData.FloatLevelSwitches, dao.FloatLevelSwitchLog{})
	newEquipmentKey(eqptKeys[3], confData.LiquidLevelGauges, dao.LiquidLevelGaugeLog{})
	newEquipmentKey(eqptKeys[4], confData.FlowMeters, dao.FlowMeterLog{})
	newEquipmentKey(eqptKeys[5], confData.WaterQualitys, dao.WaterQualityLog{})
	newEquipmentKey(eqptKeys[6], confData.Machines, dao.MachineLog{})
}

func newEquipmentKey(eqptKey string, eqptDatas, eqptStruct interface{}) {

	switch reflect.TypeOf(eqptDatas) {
	case reflect.TypeOf([]*md.KnifeGateValve{}):
		for _, data := range eqptDatas.([]*md.KnifeGateValve) {
			setEquipmentFieldDatas(eqptKey, data.Id, eqptStruct)
		}
	case reflect.TypeOf([]*md.Pump{}):
		for _, data := range eqptDatas.([]*md.Pump) {
			setEquipmentFieldDatas(eqptKey, data.Id, eqptStruct)
		}
	case reflect.TypeOf([]*md.FloatLevelSwitch{}):
		for _, data := range eqptDatas.([]*md.FloatLevelSwitch) {
			setEquipmentFieldDatas(eqptKey, data.Id, eqptStruct)
		}
	case reflect.TypeOf([]*md.LiquidLevelGauge{}):
		for _, data := range eqptDatas.([]*md.LiquidLevelGauge) {
			setEquipmentFieldDatas(eqptKey, data.Id, eqptStruct)
		}
	case reflect.TypeOf([]*md.FlowMeter{}):
		for _, data := range eqptDatas.([]*md.FlowMeter) {
			setEquipmentFieldDatas(eqptKey, data.Id, eqptStruct)
		}
	case reflect.TypeOf([]*md.WaterQuality{}):
		for _, data := range eqptDatas.([]*md.WaterQuality) {
			setEquipmentFieldDatas(eqptKey, data.Id, eqptStruct)
		}
	case reflect.TypeOf([]*md.Machine{}):
		for _, data := range eqptDatas.([]*md.Machine) {
			setEquipmentFieldDatas(eqptKey, data.Id, eqptStruct)
		}
	}
}

func setEquipmentFieldDatas(eqptKey string, id int, eqptStruct interface{}) {

	fields := make([]string, 0)

	key := util.CombineString(eqptKey, model.Colon, strconv.Itoa(id))
	columns := util.StructToDBColumns(util.GetStructFieldNames(eqptStruct))

	for _, column := range columns {
		fields = append(fields, column)
		fields = append(fields, "")
	}
	db.RedisHashSet(key, fields)
}
