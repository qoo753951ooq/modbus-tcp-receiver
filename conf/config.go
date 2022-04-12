package conf

import (
	"encoding/json"
	"fmt"
	"modbus-tcp-receiver/conf/md"
	"os"
)

var CurrentConfig Configuration

type Configuration struct {
	Param              md.Param               `json:"param"`
	KnifeGateValves    []*md.KnifeGateValve   `json:"knife_gate_valve"`
	Pumps              []*md.Pump             `json:"pump"`
	FloatLevelSwitches []*md.FloatLevelSwitch `json:"float_level_switch"`
	LiquidLevelGauges  []*md.LiquidLevelGauge `json:"liquid_level_gauge"`
	FlowMeters         []*md.FlowMeter        `json:"flow_meter"`
	WaterQualitys      []*md.WaterQuality     `json:"water_quality"`
	Machines           []*md.Machine          `json:"machine"`
}

func init() {
	file, err := os.Open("conf/config.json")

	if err != nil {
		file, _ = os.Open("./conf/config.json")
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&CurrentConfig)

	if err != nil {
		fmt.Println("Decode config error : ", err)
	}
}
