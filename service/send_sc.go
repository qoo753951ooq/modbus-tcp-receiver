package service

import (
	"modbus-tcp-receiver/conf"
	"modbus-tcp-receiver/modbus"
	"modbus-tcp-receiver/model"
	"modbus-tcp-receiver/service/eqpt"
)

type SendJob struct {
	ConfData conf.Configuration
	SendType string
}

func (this SendJob) Run() {

	err := modbus.NewModbusTcp(this.ConfData.Param.ModbusIp, false)

	if err == nil {
		sendData(this.ConfData, this.SendType)
	}

	modbus.CloseModBusTcp()
}

func sendData(confData conf.Configuration, sendType string) {
	sendKGV(confData, sendType)
	sendPump(confData, sendType)
	sendFLS(confData, sendType)
	sendLLG(confData, sendType)
	sendFM(confData, sendType)
	sendWQ(confData, sendType)
	sendMachine(confData, sendType)
}

func sendKGV(confData conf.Configuration, sendType string) {

	for _, data := range confData.KnifeGateValves {
		kgvData := eqpt.GetKnifeGeteValveData(data, sendType)

		switch sendType {
		case model.Send_List_Type:
			eqptKey := confData.Param.Redis.EquipmentKeys[0]
			go eqpt.AddKnifeGeteValveListData(eqptKey, kgvData)
		case model.Send_Log_Type:
			go eqpt.AddKnifeGeteValveLogData(kgvData)
		}
	}
}

func sendPump(confData conf.Configuration, sendType string) {

	for _, data := range confData.Pumps {
		pData := eqpt.GetPumpData(data, sendType)

		switch sendType {
		case model.Send_List_Type:
			eqptKey := confData.Param.Redis.EquipmentKeys[1]
			go eqpt.AddPumpListData(eqptKey, pData)
		case model.Send_Log_Type:
			go eqpt.AddPumpLogData(pData)
		}
	}
}

func sendFLS(confData conf.Configuration, sendType string) {

	for _, data := range confData.FloatLevelSwitches {
		flsData := eqpt.GetFloatLevelSwitchData(data, sendType)

		switch sendType {
		case model.Send_List_Type:
			eqptKey := confData.Param.Redis.EquipmentKeys[2]
			go eqpt.AddFloatLevelSwitchListData(eqptKey, flsData)
		case model.Send_Log_Type:
			go eqpt.AddFloatLevelSwitchLogData(flsData)
		}
	}
}

func sendLLG(confData conf.Configuration, sendType string) {

	for _, data := range confData.LiquidLevelGauges {
		llgData := eqpt.GetLiquidLevelGaugeData(data, sendType)

		switch sendType {
		case model.Send_List_Type:
			eqptKey := confData.Param.Redis.EquipmentKeys[3]
			go eqpt.AddLiquidLevelGaugeListData(eqptKey, llgData)
		case model.Send_Log_Type:
			go eqpt.AddLiquidLevelGaugeLogData(llgData)
		}
	}
}

func sendFM(confData conf.Configuration, sendType string) {

	for _, data := range confData.FlowMeters {
		fmData := eqpt.GetFlowMeterData(data, sendType)

		switch sendType {
		case model.Send_List_Type:
			eqptKey := confData.Param.Redis.EquipmentKeys[4]
			go eqpt.AddFlowMeterListData(eqptKey, fmData)
		case model.Send_Log_Type:
			go eqpt.AddFlowMeterLogData(fmData)
		}
	}
}

func sendWQ(confData conf.Configuration, sendType string) {

	for _, data := range confData.WaterQualitys {
		wqData := eqpt.GetWaterQualityData(data, sendType)

		switch sendType {
		case model.Send_List_Type:
			eqptKey := confData.Param.Redis.EquipmentKeys[5]
			go eqpt.AddWaterQualityListData(eqptKey, wqData)
		case model.Send_Log_Type:
			go eqpt.AddWaterQualityLogData(wqData)
		}
	}
}

func sendMachine(confData conf.Configuration, sendType string) {

	for _, data := range confData.Machines {
		mData := eqpt.GetMachineData(data, sendType)

		switch sendType {
		case model.Send_List_Type:
			eqptKey := confData.Param.Redis.EquipmentKeys[6]
			go eqpt.AddMachineListData(eqptKey, mData)
		case model.Send_Log_Type:
			go eqpt.AddMachineLogData(mData)
		}
	}
}
