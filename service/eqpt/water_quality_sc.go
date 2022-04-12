package eqpt

import (
	"modbus-tcp-receiver/conf/md"
	"modbus-tcp-receiver/modbus"
	"modbus-tcp-receiver/model"
	"modbus-tcp-receiver/model/dao"
	"modbus-tcp-receiver/util"
)

func GetWaterQualityData(data *md.WaterQuality, sendType string) dao.WaterQualityLog {

	var result dao.WaterQualityLog

	result.Id = data.Id

	wqStatus := getWaterQualityStatus(data, sendType)
	result.SsStatus = wqStatus.Ss_status
	result.DoStatus = wqStatus.Do_status
	result.PhStatus = wqStatus.Ph_status
	result.EcStatus = wqStatus.Ec_status
	result.RcStatus = wqStatus.Rc_status
	result.TpStatus = wqStatus.Tp_status

	result.Ss = getWaterQualityValueForSS(data, sendType)
	result.Do = getWaterQualityValueForDO(data, sendType)
	result.Ph = getWaterQualityValueForPH(data, sendType)
	result.Ec = getWaterQualityValueForEC(data, sendType)
	result.Rc = getWaterQualityValueForRC(data, sendType)
	result.Tp = getWaterQualityValueForTP(data, sendType)

	result.SsHighAlarm = getWaterQualityAlarmValueForSSHigh(data, sendType)
	result.SsLowAlarm = getWaterQualityAlarmValueForSSLow(data, sendType)
	result.DoHighAlarm = getWaterQualityAlarmValueForDOHigh(data, sendType)
	result.DoLowAlarm = getWaterQualityAlarmValueForDOLow(data, sendType)
	result.PhHighAlarm = getWaterQualityAlarmValueForPHHigh(data, sendType)
	result.PhLowAlarm = getWaterQualityAlarmValueForPHLow(data, sendType)
	result.EcHighAlarm = getWaterQualityAlarmValueForECHigh(data, sendType)
	result.EcLowAlarm = getWaterQualityAlarmValueForECLow(data, sendType)
	result.RcHighAlarm = getWaterQualityAlarmValueForRCHigh(data, sendType)
	result.RcLowAlarm = getWaterQualityAlarmValueForRCLow(data, sendType)
	result.TpHighAlarm = getWaterQualityAlarmValueForTPHigh(data, sendType)
	result.TpLowAlarm = getWaterQualityAlarmValueForTPLow(data, sendType)

	result.Datatime = util.GetLocationTime(util.GetDataTimeSecondToZero(util.GetTimeNow()))

	return result
}

func getWaterQualityStatus(data *md.WaterQuality, sendType string) dao.WaterQualityStatus {

	var statusData *md.Status

	for _, status := range data.MultiStatus {

		if status.Name == "status" {
			statusData = status
			break
		}
	}

	if statusData == nil {
		return dao.WaterQualityStatus{}
	}

	statusArray, _ := modbus.GetReadHoldingRegisters(uint16(statusData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(statusData.Address), modbus.Show_binary_log, statusArray)
	}

	if len(statusArray) == 0 {
		return dao.WaterQualityStatus{}
	}

	var result dao.WaterQualityStatus

	result.Ss_status = getWaterQualityStatusForSS(statusArray, statusData.Datas)
	result.Do_status = getWaterQualityStatusForDO(statusArray, statusData.Datas)
	result.Ph_status = getWaterQualityStatusForPH(statusArray, statusData.Datas)
	result.Ec_status = getWaterQualityStatusForEC(statusArray, statusData.Datas)
	result.Rc_status = getWaterQualityStatusForRC(statusArray, statusData.Datas)
	result.Tp_status = getWaterQualityStatusForTP(statusArray, statusData.Datas)

	return result
}

//取得水質SS(懸浮固體)狀態
func getWaterQualityStatusForSS(statusArray []int, statusDatas []*md.StatusData) string {
	var highIndex, lowIndex, alarmMode int

	for _, sData := range statusDatas {
		switch sData.Name {
		case "ss-high-alarm":
			highIndex = util.BitIndex2ArrayIndex(sData.Index)
			alarmMode = alarmMode + 1
		case "ss-low-alarm":
			lowIndex = util.BitIndex2ArrayIndex(sData.Index)
			alarmMode = alarmMode + 2
		}
	}

	if highIndex == -1 || lowIndex == -1 {
		return model.Not_Equipment_Field_Status
	}

	ssStatus := getWQSensorAlarmStatus(statusArray, highIndex, lowIndex, alarmMode)
	return ssStatus
}

//取得水質DO(溶氧)狀態
func getWaterQualityStatusForDO(statusArray []int, statusDatas []*md.StatusData) string {
	var highIndex, lowIndex, alarmMode int

	for _, sData := range statusDatas {
		switch sData.Name {
		case "do-high-alarm":
			highIndex = util.BitIndex2ArrayIndex(sData.Index)
			alarmMode = alarmMode + 1
		case "do-low-alarm":
			lowIndex = util.BitIndex2ArrayIndex(sData.Index)
			alarmMode = alarmMode + 2
		}
	}

	if highIndex == -1 || lowIndex == -1 {
		return model.Not_Equipment_Field_Status
	}

	doStatus := getWQSensorAlarmStatus(statusArray, highIndex, lowIndex, alarmMode)
	return doStatus
}

//取得水質PH(酸鹼度)狀態
func getWaterQualityStatusForPH(statusArray []int, statusDatas []*md.StatusData) string {
	var highIndex, lowIndex, alarmMode int

	for _, sData := range statusDatas {
		switch sData.Name {
		case "ph-high-alarm":
			highIndex = util.BitIndex2ArrayIndex(sData.Index)
			alarmMode = alarmMode + 1
		case "ph-low-alarm":
			lowIndex = util.BitIndex2ArrayIndex(sData.Index)
			alarmMode = alarmMode + 2
		}
	}

	if highIndex == -1 || lowIndex == -1 {
		return model.Not_Equipment_Field_Status
	}

	phStatus := getWQSensorAlarmStatus(statusArray, highIndex, lowIndex, alarmMode)

	return phStatus
}

//取得水質EC(導電度)狀態
func getWaterQualityStatusForEC(statusArray []int, statusDatas []*md.StatusData) string {
	var highIndex, lowIndex, alarmMode int

	for _, sData := range statusDatas {
		switch sData.Name {
		case "ec-high-alarm":
			highIndex = util.BitIndex2ArrayIndex(sData.Index)
			alarmMode = alarmMode + 1
		case "ec-low-alarm":
			lowIndex = util.BitIndex2ArrayIndex(sData.Index)
			alarmMode = alarmMode + 2
		}

	}

	if highIndex == -1 || lowIndex == -1 {
		return model.Not_Equipment_Field_Status
	}

	ecStatus := getWQSensorAlarmStatus(statusArray, highIndex, lowIndex, alarmMode)
	return ecStatus
}

//取得水質RC(餘氯)狀態
func getWaterQualityStatusForRC(statusArray []int, statusDatas []*md.StatusData) string {
	var highIndex, lowIndex, alarmMode int

	for _, sData := range statusDatas {
		switch sData.Name {
		case "rc-high-alarm":
			highIndex = util.BitIndex2ArrayIndex(sData.Index)
			alarmMode = alarmMode + 1
		case "rc-low-alarm":
			lowIndex = util.BitIndex2ArrayIndex(sData.Index)
			alarmMode = alarmMode + 2
		}

	}

	if highIndex == -1 || lowIndex == -1 {
		return model.Not_Equipment_Field_Status
	}

	rcStatus := getWQSensorAlarmStatus(statusArray, highIndex, lowIndex, alarmMode)
	return rcStatus
}

//取得水質Tp(溫度)狀態
func getWaterQualityStatusForTP(statusArray []int, statusDatas []*md.StatusData) string {
	var highIndex, lowIndex, alarmMode int

	for _, sData := range statusDatas {
		switch sData.Name {
		case "tp-high-alarm":
			highIndex = util.BitIndex2ArrayIndex(sData.Index)
			alarmMode = alarmMode + 1
		case "tp-low-alarm":
			lowIndex = util.BitIndex2ArrayIndex(sData.Index)
			alarmMode = alarmMode + 2
		}

	}

	if highIndex == -1 || lowIndex == -1 {
		return model.Not_Equipment_Field_Status
	}

	tpStatus := getWQSensorAlarmStatus(statusArray, highIndex, lowIndex, alarmMode)
	return tpStatus
}

/**取得水質感測器警報狀態結果
(1) 高
(2) 低
(3) 高+低*/
func getWQSensorAlarmStatus(statusArray []int, highIndex, lowIndex, alarmMode int) string {

	var alarmStatus string

	switch {
	case alarmMode == 1:
		switch {
		case statusArray[highIndex] == 0:
			alarmStatus = model.Status_Normal
		case statusArray[highIndex] == 1:
			alarmStatus = model.Status_High_Alarm
		}
	case alarmMode == 2:
		switch {
		case statusArray[lowIndex] == 0:
			alarmStatus = model.Status_Normal
		case statusArray[lowIndex] == 1:
			alarmStatus = model.Status_Low_Alarm
		}
	case alarmMode == 3:
		switch {
		case statusArray[highIndex] == 0 && statusArray[lowIndex] == 0:
			alarmStatus = model.Status_Normal
		case statusArray[highIndex] == 1 && statusArray[lowIndex] == 0:
			alarmStatus = model.Status_High_Alarm
		case statusArray[highIndex] == 0 && statusArray[lowIndex] == 1:
			alarmStatus = model.Status_Low_Alarm
		}
	}

	return alarmStatus
}

//取得水質SS(懸浮固體)值
func getWaterQualityValueForSS(data *md.WaterQuality, sendType string) float64 {
	var ssData *md.Value

	for _, value := range data.Values {
		if value.Name == "ss" {
			ssData = value
			break
		}
	}

	if ssData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, ssArray := modbus.GetReadHoldingRegisters(uint16(ssData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(ssData.Address), modbus.Show_decimal_log, ssArray)
	}

	if len(ssArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(ssArray[0], ssData.Decimal_places)
}

//取得水質DO(溶氧)值
func getWaterQualityValueForDO(data *md.WaterQuality, sendType string) float64 {
	var doData *md.Value

	for _, value := range data.Values {
		if value.Name == "do" {
			doData = value
			break
		}
	}

	if doData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, doArray := modbus.GetReadHoldingRegisters(uint16(doData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(doData.Address), modbus.Show_decimal_log, doArray)
	}

	if len(doArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(doArray[0], doData.Decimal_places)
}

//取得水質PH(酸鹼度)值
func getWaterQualityValueForPH(data *md.WaterQuality, sendType string) float64 {
	var phData *md.Value

	for _, value := range data.Values {
		if value.Name == "ph" {
			phData = value
			break
		}
	}

	if phData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, phArray := modbus.GetReadHoldingRegisters(uint16(phData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(phData.Address), modbus.Show_decimal_log, phArray)
	}

	if len(phArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(phArray[0], phData.Decimal_places)
}

//取得水質EC(導電度)值
func getWaterQualityValueForEC(data *md.WaterQuality, sendType string) float64 {
	var ecData *md.Value

	for _, value := range data.Values {
		if value.Name == "ec" {
			ecData = value
			break
		}
	}

	if ecData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, ecArray := modbus.GetReadHoldingRegisters(uint16(ecData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(ecData.Address), modbus.Show_decimal_log, ecArray)
	}

	if len(ecArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(ecArray[0], ecData.Decimal_places)
}

//取得水質RC(餘氯)值
func getWaterQualityValueForRC(data *md.WaterQuality, sendType string) float64 {
	var rcData *md.Value

	for _, value := range data.Values {
		if value.Name == "rc" {
			rcData = value
			break
		}
	}

	if rcData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, rcArray := modbus.GetReadHoldingRegisters(uint16(rcData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(rcData.Address), modbus.Show_decimal_log, rcArray)
	}

	if len(rcArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(rcArray[0], rcData.Decimal_places)
}

//取得水質TP(溫度)值
func getWaterQualityValueForTP(data *md.WaterQuality, sendType string) float64 {
	var tpData *md.Value

	for _, value := range data.Values {
		if value.Name == "tp" {
			tpData = value
			break
		}
	}

	if tpData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, tpArray := modbus.GetReadHoldingRegisters(uint16(tpData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(tpData.Address), modbus.Show_decimal_log, tpArray)
	}

	if len(tpArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(tpArray[0], tpData.Decimal_places)
}

//取得水質SS(懸浮固體)高警報值
func getWaterQualityAlarmValueForSSHigh(data *md.WaterQuality, sendType string) float64 {
	var ssHData *md.Value

	for _, value := range data.Values {
		if value.Name == "ss-high-alarm" {
			ssHData = value
			break
		}
	}

	if ssHData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, ssHArray := modbus.GetReadHoldingRegisters(uint16(ssHData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(ssHData.Address), modbus.Show_decimal_log, ssHArray)
	}

	if len(ssHArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(ssHArray[0], ssHData.Decimal_places)
}

//取得水質SS(懸浮固體)低警報值
func getWaterQualityAlarmValueForSSLow(data *md.WaterQuality, sendType string) float64 {
	var ssLData *md.Value

	for _, value := range data.Values {
		if value.Name == "ss-low-alarm" {
			ssLData = value
			break
		}
	}

	if ssLData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, ssLArray := modbus.GetReadHoldingRegisters(uint16(ssLData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(ssLData.Address), modbus.Show_decimal_log, ssLArray)
	}

	if len(ssLArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(ssLArray[0], ssLData.Decimal_places)
}

//取得水質DO(溶氧)高警報值
func getWaterQualityAlarmValueForDOHigh(data *md.WaterQuality, sendType string) float64 {
	var doHData *md.Value

	for _, value := range data.Values {
		if value.Name == "do-high-alarm" {
			doHData = value
			break
		}
	}

	if doHData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, doHArray := modbus.GetReadHoldingRegisters(uint16(doHData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(doHData.Address), modbus.Show_decimal_log, doHArray)
	}

	if len(doHArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(doHArray[0], doHData.Decimal_places)
}

//取得水質DO(溶氧)低警報值
func getWaterQualityAlarmValueForDOLow(data *md.WaterQuality, sendType string) float64 {
	var doLData *md.Value

	for _, value := range data.Values {
		if value.Name == "do-low-alarm" {
			doLData = value
			break
		}
	}

	if doLData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, doLArray := modbus.GetReadHoldingRegisters(uint16(doLData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(doLData.Address), modbus.Show_decimal_log, doLArray)
	}

	if len(doLArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(doLArray[0], doLData.Decimal_places)
}

//取得水質PH(酸鹼度)高警報值
func getWaterQualityAlarmValueForPHHigh(data *md.WaterQuality, sendType string) float64 {
	var phHData *md.Value

	for _, value := range data.Values {
		if value.Name == "ph-high-alarm" {
			phHData = value
			break
		}
	}

	if phHData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, phHArray := modbus.GetReadHoldingRegisters(uint16(phHData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(phHData.Address), modbus.Show_decimal_log, phHArray)
	}

	if len(phHArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(phHArray[0], phHData.Decimal_places)
}

//取得水質PH(酸鹼度)低警報值
func getWaterQualityAlarmValueForPHLow(data *md.WaterQuality, sendType string) float64 {
	var phLData *md.Value

	for _, value := range data.Values {
		if value.Name == "ph-low-alarm" {
			phLData = value
			break
		}
	}

	if phLData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, phLArray := modbus.GetReadHoldingRegisters(uint16(phLData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(phLData.Address), modbus.Show_decimal_log, phLArray)
	}

	if len(phLArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(phLArray[0], phLData.Decimal_places)
}

//取得水質EC(導電度)高警報值
func getWaterQualityAlarmValueForECHigh(data *md.WaterQuality, sendType string) float64 {
	var ecHData *md.Value

	for _, value := range data.Values {
		if value.Name == "ec-high-alarm" {
			ecHData = value
			break
		}
	}

	if ecHData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, ecHArray := modbus.GetReadHoldingRegisters(uint16(ecHData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(ecHData.Address), modbus.Show_decimal_log, ecHArray)
	}

	if len(ecHArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(ecHArray[0], ecHData.Decimal_places)
}

//取得水質EC(導電度)低警報值
func getWaterQualityAlarmValueForECLow(data *md.WaterQuality, sendType string) float64 {
	var ecLData *md.Value

	for _, value := range data.Values {
		if value.Name == "ec-low-alarm" {
			ecLData = value
			break
		}
	}

	if ecLData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, ecLArray := modbus.GetReadHoldingRegisters(uint16(ecLData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(ecLData.Address), modbus.Show_decimal_log, ecLArray)
	}

	if len(ecLArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(ecLArray[0], ecLData.Decimal_places)
}

//取得水質RC(餘氯)高警報值
func getWaterQualityAlarmValueForRCHigh(data *md.WaterQuality, sendType string) float64 {
	var rcHData *md.Value

	for _, value := range data.Values {
		if value.Name == "rc-high-alarm" {
			rcHData = value
			break
		}
	}

	if rcHData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, rcHArray := modbus.GetReadHoldingRegisters(uint16(rcHData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(rcHData.Address), modbus.Show_decimal_log, rcHArray)
	}

	if len(rcHArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(rcHArray[0], rcHData.Decimal_places)
}

//取得水質RC(餘氯)低警報值
func getWaterQualityAlarmValueForRCLow(data *md.WaterQuality, sendType string) float64 {
	var rcLData *md.Value

	for _, value := range data.Values {
		if value.Name == "rc-low-alarm" {
			rcLData = value
			break
		}
	}

	if rcLData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, rcLArray := modbus.GetReadHoldingRegisters(uint16(rcLData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(rcLData.Address), modbus.Show_decimal_log, rcLArray)
	}

	if len(rcLArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(rcLArray[0], rcLData.Decimal_places)
}

//取得水質TP(溫度)高警報值
func getWaterQualityAlarmValueForTPHigh(data *md.WaterQuality, sendType string) float64 {
	var tpHData *md.Value

	for _, value := range data.Values {
		if value.Name == "tp-high-alarm" {
			tpHData = value
			break
		}
	}

	if tpHData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, tpHArray := modbus.GetReadHoldingRegisters(uint16(tpHData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(tpHData.Address), modbus.Show_decimal_log, tpHArray)
	}

	if len(tpHArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(tpHArray[0], tpHData.Decimal_places)
}

//取得水質TP(溫度)低警報值
func getWaterQualityAlarmValueForTPLow(data *md.WaterQuality, sendType string) float64 {
	var tpLData *md.Value

	for _, value := range data.Values {
		if value.Name == "tp-low-alarm" {
			tpLData = value
			break
		}
	}

	if tpLData == nil {
		return model.Not_Equipment_Field_Value
	}

	_, tpLArray := modbus.GetReadHoldingRegisters(uint16(tpLData.Address), 1)

	if stype := isLogSendType(sendType); stype == true {
		modbus.ShowModBusLog(uint16(tpLData.Address), modbus.Show_decimal_log, tpLArray)
	}

	if len(tpLArray) == 0 {
		return model.Not_Equipment_Field_Value
	}

	return util.MultiplyDP(tpLArray[0], tpLData.Decimal_places)
}
