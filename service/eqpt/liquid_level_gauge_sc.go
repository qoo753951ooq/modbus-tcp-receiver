package eqpt

import (
	"errors"
	"modbus-tcp-receiver/conf/md"
	"modbus-tcp-receiver/modbus"
	"modbus-tcp-receiver/model"
	"modbus-tcp-receiver/model/dao"
	"modbus-tcp-receiver/util"
	"strconv"
)

func GetLiquidLevelGaugeData(data *md.LiquidLevelGauge, sendType string) dao.LiquidLevelGaugeLog {

	var result dao.LiquidLevelGaugeLog

	result.Id = data.Id

	result.Level = getLiquidLevelGaugeStatusForLevel(data, sendType)
	result.Status = getLiquidLevelGaugeStatusForAlarm(data, sendType)
	result.Liquid = getLiquidLevelGaugeLiquidValue(data, sendType)
	result.LiquidHighAlarm = getLiquidLevelGaugeAlarmValueForLiquidHigh(data, sendType)
	result.LiquidLowAlarm = getLiquidLevelGaugeAlarmValueForLiquidLow(data, sendType)
	result.Datatime = util.GetLocationTime(util.GetDataTimeSecondToZero(util.GetTimeNow()))

	return result
}

func AddLiquidLevelGaugeListData(eqptKey string, llgData dao.LiquidLevelGaugeLog) error {

	util.AddLiquidLevelGaugeList.Lock()

	err := llgData.UpdateLiquidLevelGaugeList(eqptKey)

	util.AddLiquidLevelGaugeList.Unlock()

	if err != nil {
		return errors.New(util.CombineString(strconv.Itoa(llgData.Id), " redis set failure"))
	}

	return nil
}

func AddLiquidLevelGaugeLogData(llgData dao.LiquidLevelGaugeLog) error {

	util.AddLiquidLevelGaugeLog.Lock()

	_, err := llgData.InsertLiquidLevelGaugeLog()

	util.AddLiquidLevelGaugeLog.Unlock()

	if err != nil {
		return errors.New(util.CombineString(strconv.Itoa(llgData.Id), " pg insert failure"))
	}
	return nil
}

//取得液位計level(水位高度)狀態
func getLiquidLevelGaugeStatusForLevel(data *md.LiquidLevelGauge, sendType string) string {
	var levelData *md.Status

	for _, status := range data.MultiStatus {

		if status.Name == "level" {
			levelData = status
			break
		}
	}

	if levelData == nil {
		return model.Not_Equipment_Field_Status
	}

	levelArray, _ := modbus.GetReadHoldingRegisters(uint16(levelData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(levelData.Address), modbus.Show_binary_log, levelArray)
	}

	if len(levelArray) == 0 {
		return model.Not_Equipment_Field_Status
	}

	var llIndex, lIndex, mIndex, hIndex, hhIndex int
	var levelStatus string

	for _, lData := range levelData.Datas {

		switch lData.Name {
		case "ll-level":
			llIndex = util.BitIndex2ArrayIndex(lData.Index)
		case "l-level":
			lIndex = util.BitIndex2ArrayIndex(lData.Index)
		case "m-level":
			mIndex = util.BitIndex2ArrayIndex(lData.Index)
		case "h-level":
			hIndex = util.BitIndex2ArrayIndex(lData.Index)
		case "hh-level":
			hhIndex = util.BitIndex2ArrayIndex(lData.Index)
		}
	}

	if llIndex == -1 || lIndex == -1 || mIndex == -1 || hIndex == -1 || hhIndex == -1 {
		return model.Not_Equipment_Field_Status
	}

	levelList := getLevelList(data.Level_mode,
		levelArray[llIndex], levelArray[lIndex], levelArray[mIndex], levelArray[hIndex], levelArray[hhIndex])

	var levelStatusCount int

	for _, lv := range levelList {

		if lv.LevelValue == 1 {
			levelStatusCount++
		}
	}

	if levelStatusCount == 0 {
		return model.Not_Equipment_Field_Status
	}

	levelStatus = levelList[levelStatusCount-1].LevelStatus
	return levelStatus
}

//取得液位計alarm(水位警報)狀態
func getLiquidLevelGaugeStatusForAlarm(data *md.LiquidLevelGauge, sendType string) string {
	var statusData *md.Status

	for _, status := range data.MultiStatus {

		if status.Name == "alarm" {
			statusData = status
			break
		}
	}

	if statusData == nil {
		return model.Not_Equipment_Field_Status
	}

	statusArray, _ := modbus.GetReadHoldingRegisters(uint16(statusData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(statusData.Address), modbus.Show_binary_log, statusArray)
	}

	if len(statusArray) == 0 {
		return model.Not_Equipment_Field_Status
	}

	var llIndex, hhIndex int
	var status string

	for _, sData := range statusData.Datas {

		switch sData.Name {
		case "ll-alarm":
			llIndex = util.BitIndex2ArrayIndex(sData.Index)
		case "hh-alarm":
			hhIndex = util.BitIndex2ArrayIndex(sData.Index)
		}
	}

	if llIndex == -1 || hhIndex == -1 {
		return model.Not_Equipment_Field_Status
	}

	switch {
	case statusArray[llIndex] == 0 && statusArray[hhIndex] == 0:
		status = model.Status_Normal
	case statusArray[llIndex] == 1 && statusArray[hhIndex] == 0:
		status = model.Status_Low_Alarm
	case statusArray[llIndex] == 0 && statusArray[hhIndex] == 1:
		status = model.Status_High_Alarm
	}

	return status
}

//取得液位計Liquid(液位)值
func getLiquidLevelGaugeLiquidValue(data *md.LiquidLevelGauge, sendType string) float64 {

	var liquidData *md.Value

	for _, value := range data.Values {

		if value.Name == "liquid" {
			liquidData = value
			break
		}
	}

	if liquidData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, liquidArray := modbus.GetReadHoldingRegisters(uint16(liquidData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(liquidData.Address), modbus.Show_decimal_log, liquidArray)
	}

	if len(liquidArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(liquidArray[0], liquidData.Decimal_places)
}

//取得液位計Liquid(液位)高警報值
func getLiquidLevelGaugeAlarmValueForLiquidHigh(data *md.LiquidLevelGauge, sendType string) float64 {
	var liquidHData *md.Value

	for _, value := range data.Values {
		if value.Name == "liquid-high-alarm" {
			liquidHData = value
			break
		}
	}

	if liquidHData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, liquidHArray := modbus.GetReadHoldingRegisters(uint16(liquidHData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(liquidHData.Address), modbus.Show_decimal_log, liquidHArray)
	}

	if len(liquidHArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(liquidHArray[0], liquidHData.Decimal_places)
}

//取得液位計Liquid(液位)低警報值
func getLiquidLevelGaugeAlarmValueForLiquidLow(data *md.LiquidLevelGauge, sendType string) float64 {
	var liquidLData *md.Value

	for _, value := range data.Values {
		if value.Name == "liquid-low-alarm" {
			liquidLData = value
			break
		}
	}

	if liquidLData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, liquidLArray := modbus.GetReadHoldingRegisters(uint16(liquidLData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(liquidLData.Address), modbus.Show_decimal_log, liquidLArray)
	}

	if len(liquidLArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(liquidLArray[0], liquidLData.Decimal_places)
}
