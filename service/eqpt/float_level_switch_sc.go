package eqpt

import (
	"modbus-tcp-receiver/conf/md"
	"modbus-tcp-receiver/modbus"
	"modbus-tcp-receiver/model"
	"modbus-tcp-receiver/model/dao"
	"modbus-tcp-receiver/util"
)

func GetFloatLevelSwitchData(data *md.FloatLevelSwitch, sendType string) dao.FloatLevelSwitchLog {

	var result dao.FloatLevelSwitchLog

	result.Id = data.Id

	result.Level = getFloatLevelSwitchStatusForLevel(data, sendType)
	result.Status = getFloatLevelSwitchStatusForAlarm(data, sendType)
	result.Datatime = util.GetLocationTime(util.GetDataTimeSecondToZero(util.GetTimeNow()))

	return result
}

//取得液位開關level(水位高度)狀態
func getFloatLevelSwitchStatusForLevel(data *md.FloatLevelSwitch, sendType string) string {
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

//取得液位開關alarm(水位警報)狀態
func getFloatLevelSwitchStatusForAlarm(data *md.FloatLevelSwitch, sendType string) string {
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
