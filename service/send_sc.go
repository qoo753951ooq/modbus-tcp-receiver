package service

import (
	"fmt"
	"modbus-tcp-receiver/conf"
	"modbus-tcp-receiver/modbus"
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
		fmt.Printf("%v \n", kgvData)
	}
}

func sendPump(confData conf.Configuration, sendType string) {

	for _, data := range confData.Pumps {
		pData := eqpt.GetPumpData(data, sendType)
		fmt.Printf("%v \n", pData)
	}
}

func sendFLS(confData conf.Configuration, sendType string) {

	for _, data := range confData.FloatLevelSwitches {
		flsData := eqpt.GetFloatLevelSwitchData(data, sendType)
		fmt.Printf("%v \n", flsData)
	}
}

func sendLLG(confData conf.Configuration, sendType string) {

	for _, data := range confData.LiquidLevelGauges {
		llgData := eqpt.GetLiquidLevelGaugeData(data, sendType)
		fmt.Printf("%v \n", llgData)
	}
}

func sendFM(confData conf.Configuration, sendType string) {

	for _, data := range confData.FlowMeters {
		fmData := eqpt.GetFlowMeterData(data, sendType)
		fmt.Printf("%v \n", fmData)
	}
}

func sendWQ(confData conf.Configuration, sendType string) {

	for _, data := range confData.WaterQualitys {
		wqData := eqpt.GetWaterQualityData(data, sendType)
		fmt.Printf("%v \n", wqData)
	}
}

func sendMachine(confData conf.Configuration, sendType string) {

	for _, data := range confData.Machines {
		mData := eqpt.GetMachineData(data, sendType)
		fmt.Printf("%v \n", mData)
	}
}
