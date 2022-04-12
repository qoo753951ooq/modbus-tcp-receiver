package eqpt

import (
	"modbus-tcp-receiver/conf/md"
	"modbus-tcp-receiver/modbus"
	"modbus-tcp-receiver/model"
	"modbus-tcp-receiver/model/dao"
	"modbus-tcp-receiver/util"
)

func GetKnifeGeteValveData(data *md.KnifeGateValve, sendType string) dao.KnifeGateValveLog {

	var result dao.KnifeGateValveLog

	result.Id = data.Id
	result.Status = getKnifeGeteValveStatus(data, sendType)
	result.Opening = getKnifeGeteValveOpening(data, sendType)
	result.Datatime = util.GetLocationTime(util.GetDataTimeSecondToZero(util.GetTimeNow()))

	return result
}

//取得閘刀閥狀態
func getKnifeGeteValveStatus(data *md.KnifeGateValve, sendType string) string {
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

	var opIndex, clIndex, errIndex int
	var status string

	for _, sData := range statusData.Datas {

		switch sData.Name {
		case "all-open":
			opIndex = util.BitIndex2ArrayIndex(sData.Index)
		case "all-close":
			clIndex = util.BitIndex2ArrayIndex(sData.Index)
		case "error":
			errIndex = util.BitIndex2ArrayIndex(sData.Index)
		}
	}

	if opIndex == -1 || clIndex == -1 || errIndex == -1 {
		return model.Not_Equipment_Field_Status
	}

	if statusArray[errIndex] == 0 {
		status = model.Status_Normal
	}

	if statusArray[errIndex] == 1 {
		status = model.Status_Error
	}

	if statusArray[errIndex] == 0 && statusArray[opIndex] == 1 {
		status = model.Status_All_Open
	}

	if statusArray[errIndex] == 0 && statusArray[clIndex] == 1 {
		status = model.Status_All_Close
	}

	return status
}

//取得閘刀閥開度
func getKnifeGeteValveOpening(data *md.KnifeGateValve, sendType string) float64 {

	var openingData *md.Value

	for _, value := range data.Values {
		if value.Name == "opening" {
			openingData = value
			break
		}
	}

	if openingData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, openingArray := modbus.GetReadHoldingRegisters(uint16(openingData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(openingData.Address), modbus.Show_decimal_log, openingArray)
	}

	if len(openingArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(openingArray[0], openingData.Decimal_places)
}
