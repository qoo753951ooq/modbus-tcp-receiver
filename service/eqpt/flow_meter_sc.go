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

func GetFlowMeterData(data *md.FlowMeter, sendType string) dao.FlowMeterLog {

	var result dao.FlowMeterLog

	result.Id = data.Id
	result.FlowStatus = getFlowMeterFlowStatus(data, sendType)
	result.Flow = getFlowMeterFlowValue(data, sendType)
	result.FlowLowAlarm = getFlowMeterFlowLowAlarmValue(data, sendType)
	result.Datatime = util.GetLocationTime(util.GetDataTimeSecondToZero(util.GetTimeNow()))

	return result
}

func AddFlowMeterListData(eqptKey string, fmData dao.FlowMeterLog) error {

	util.AddFlowMeterList.Lock()

	err := fmData.UpdateFlowMeterList(eqptKey)

	util.AddFlowMeterList.Unlock()

	if err != nil {
		return errors.New(util.CombineString(strconv.Itoa(fmData.Id), " redis set failure"))
	}

	return nil
}

func AddFlowMeterLogData(fmData dao.FlowMeterLog) error {

	util.AddFlowMeterLog.Lock()

	_, err := fmData.InsertFlowMeterLog()

	util.AddFlowMeterLog.Unlock()

	if err != nil {
		return errors.New(util.CombineString(strconv.Itoa(fmData.Id), " pg insert failure"))
	}
	return nil
}

//取得流量計流量狀態
func getFlowMeterFlowStatus(data *md.FlowMeter, sendType string) string {
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

	var lowIndex int
	var flowStatus string

	for _, sData := range statusData.Datas {

		switch sData.Name {
		case "flow-low-alarm":
			lowIndex = util.BitIndex2ArrayIndex(sData.Index)
		}
	}

	if lowIndex == -1 {
		return model.Not_Equipment_Field_Status
	}

	switch {
	case statusArray[lowIndex] == 0:
		flowStatus = model.Status_Normal
	case statusArray[lowIndex] == 1:
		flowStatus = model.Status_Low_Alarm
	}

	return flowStatus
}

//取得流量計流量值
func getFlowMeterFlowValue(data *md.FlowMeter, sendType string) float64 {

	var flowData *md.Value

	for _, value := range data.Values {
		if value.Name == "flow" {
			flowData = value
			break
		}
	}

	if flowData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, flowArray := modbus.GetReadHoldingRegisters(uint16(flowData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(flowData.Address), modbus.Show_decimal_log, flowArray)
	}

	if len(flowArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(flowArray[0], flowData.Decimal_places)
}

//取得流量計低流量警報值
func getFlowMeterFlowLowAlarmValue(data *md.FlowMeter, sendType string) float64 {

	var flowLowAlarmData *md.Value

	for _, value := range data.Values {
		if value.Name == "flow-low-alarm" {
			flowLowAlarmData = value
			break
		}
	}

	if flowLowAlarmData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, flowLowAlarmArray := modbus.GetReadHoldingRegisters(uint16(flowLowAlarmData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(flowLowAlarmData.Address), modbus.Show_decimal_log, flowLowAlarmArray)
	}

	if len(flowLowAlarmArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(flowLowAlarmArray[0], flowLowAlarmData.Decimal_places)
}
