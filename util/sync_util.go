package util

import (
	"sync"
)

var AddKnifeGeteValveLog sync.Mutex
var AddKnifeGeteValveList sync.Mutex
var AddPumpLog sync.Mutex
var AddPumpList sync.Mutex
var AddFloatLevelSwitchLog sync.Mutex
var AddFloatLevelSwitchList sync.Mutex
var AddLiquidLevelGaugeLog sync.Mutex
var AddLiquidLevelGaugeList sync.Mutex
var AddFlowMeterLog sync.Mutex
var AddFlowMeterList sync.Mutex
var AddWaterQualityLog sync.Mutex
var AddWaterQualityList sync.Mutex
var AddMachineLog sync.Mutex
var AddMachineList sync.Mutex
