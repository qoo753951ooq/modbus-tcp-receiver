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

func GetMachineData(data *md.Machine, sendType string) dao.MachineLog {

	var result dao.MachineLog

	result.Id = data.Id
	result.Status = getMachineStatus(data, sendType)
	result.Datatime = util.GetLocationTime(util.GetDataTimeSecondToZero(util.GetTimeNow()))

	return result
}

func AddMachineListData(eqptKey string, mData dao.MachineLog) error {

	util.AddMachineList.Lock()

	err := mData.UpdateMachineList(eqptKey)

	util.AddMachineList.Unlock()

	if err != nil {
		return errors.New(util.CombineString(strconv.Itoa(mData.Id), " redis set failure"))
	}

	return nil
}

func AddMachineLogData(mData dao.MachineLog) error {

	util.AddMachineLog.Lock()

	_, err := mData.InsertMachineLog()

	util.AddMachineLog.Unlock()

	if err != nil {
		return errors.New(util.CombineString(strconv.Itoa(mData.Id), " pg insert failure"))
	}

	return nil
}

//取得機器狀態
func getMachineStatus(data *md.Machine, sendType string) string {
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
