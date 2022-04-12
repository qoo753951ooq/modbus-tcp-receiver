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
}

func sendKGV(confData conf.Configuration, sendType string) {

	for _, data := range confData.KnifeGateValves {

		kgvData := eqpt.GetKnifeGeteValveData(data, sendType)
		fmt.Printf("%v \n", kgvData)
	}
}
