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

func GetPumpData(data *md.Pump, sendType string) dao.PumpLog {

	var result dao.PumpLog

	result.Id = data.Id
	result.Status = getPumpStatus(data, sendType)
	result.Datatime = util.GetLocationTime(util.GetDataTimeSecondToZero(util.GetTimeNow()))

	return result
}

func AddPumpListData(eqptKey string, pData dao.PumpLog) error {

	util.AddPumpList.Lock()

	err := pData.UpdatePumpList(eqptKey)

	util.AddPumpList.Unlock()

	if err != nil {
		return errors.New(util.CombineString(strconv.Itoa(pData.Id), " redis set failure"))
	}

	return nil
}

func AddPumpLogData(pData dao.PumpLog) error {

	util.AddPumpLog.Lock()

	_, err := pData.InsertPumpLog()

	util.AddPumpLog.Unlock()

	if err != nil {
		return errors.New(util.CombineString(strconv.Itoa(pData.Id), " pg insert failure"))
	}
	return nil
}

//取得泵浦狀態
func getPumpStatus(data *md.Pump, sendType string) string {
	var statusData *md.Status

	for _, status := range data.MultiStatus {

		if status.Name == "status" {
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

	var startIndex, errIndex int
	var status string

	for _, sData := range statusData.Datas {

		switch sData.Name {
		case "start":
			startIndex = util.BitIndex2ArrayIndex(sData.Index)
		case "error":
			errIndex = util.BitIndex2ArrayIndex(sData.Index)
		}
	}

	if startIndex == -1 || errIndex == -1 {
		return model.Not_Equipment_Field_Status
	}

	if statusArray[errIndex] == 1 {
		status = model.Status_Error
	}

	if statusArray[errIndex] == 0 && statusArray[startIndex] == 1 {
		status = model.Status_Start_Up
	}

	if statusArray[errIndex] == 0 && statusArray[startIndex] == 0 {
		status = model.Status_Stop
	}

	return status
}
