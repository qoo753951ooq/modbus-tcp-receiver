package eqpt

import "modbus-tcp-receiver/model"

func isLogSendType(sendType string) bool {

	var result bool

	switch sendType {
	case model.Send_Log_Type:
		result = true
	case model.Send_List_Type:
		result = false
	}

	return result
}

/**
液位開關與液位計的Level
(1) LL,L,M,H,HH
(2) LL,L,M,H
(3) LL,L,H,HH
(4) L,H,HH
(5) L,M,H
(6)	L,H
(7) LL,L
(8)	HH		*/
func getLevelList(levelMode string, llValue, lValue, mValue, hValue, hhValue int) []*model.Level {

	result := make([]*model.Level, 0)

	switch levelMode {
	case "LL,L,M,H,HH":
		llData := model.Level{LevelStatus: model.Status_LL_Level, LevelValue: llValue}
		lData := model.Level{LevelStatus: model.Status_L_Level, LevelValue: lValue}
		mData := model.Level{LevelStatus: model.Status_M_Level, LevelValue: mValue}
		hData := model.Level{LevelStatus: model.Status_H_Level, LevelValue: hValue}
		hhData := model.Level{LevelStatus: model.Status_HH_Level, LevelValue: hhValue}
		result = append(result, &llData)
		result = append(result, &lData)
		result = append(result, &mData)
		result = append(result, &hData)
		result = append(result, &hhData)
	case "LL,L,M,H":
		llData := model.Level{LevelStatus: model.Status_LL_Level, LevelValue: llValue}
		lData := model.Level{LevelStatus: model.Status_L_Level, LevelValue: lValue}
		mData := model.Level{LevelStatus: model.Status_M_Level, LevelValue: mValue}
		hData := model.Level{LevelStatus: model.Status_H_Level, LevelValue: hValue}
		result = append(result, &llData)
		result = append(result, &lData)
		result = append(result, &mData)
		result = append(result, &hData)
	case "LL,L,H,HH":
		llData := model.Level{LevelStatus: model.Status_LL_Level, LevelValue: llValue}
		lData := model.Level{LevelStatus: model.Status_L_Level, LevelValue: lValue}
		hData := model.Level{LevelStatus: model.Status_H_Level, LevelValue: hValue}
		hhData := model.Level{LevelStatus: model.Status_HH_Level, LevelValue: hhValue}
		result = append(result, &llData)
		result = append(result, &lData)
		result = append(result, &hData)
		result = append(result, &hhData)
	case "L,H,HH":
		lData := model.Level{LevelStatus: model.Status_L_Level, LevelValue: lValue}
		hData := model.Level{LevelStatus: model.Status_H_Level, LevelValue: hValue}
		hhData := model.Level{LevelStatus: model.Status_HH_Level, LevelValue: hhValue}
		result = append(result, &lData)
		result = append(result, &hData)
		result = append(result, &hhData)
	case "L,M,H":
		lData := model.Level{LevelStatus: model.Status_L_Level, LevelValue: lValue}
		mData := model.Level{LevelStatus: model.Status_M_Level, LevelValue: mValue}
		hData := model.Level{LevelStatus: model.Status_H_Level, LevelValue: hValue}
		result = append(result, &lData)
		result = append(result, &mData)
		result = append(result, &hData)
	case "LL,L,H":
		llData := model.Level{LevelStatus: model.Status_LL_Level, LevelValue: llValue}
		lData := model.Level{LevelStatus: model.Status_L_Level, LevelValue: lValue}
		hData := model.Level{LevelStatus: model.Status_H_Level, LevelValue: hValue}
		result = append(result, &llData)
		result = append(result, &lData)
		result = append(result, &hData)
	case "L,H":
		lData := model.Level{LevelStatus: model.Status_L_Level, LevelValue: lValue}
		hData := model.Level{LevelStatus: model.Status_H_Level, LevelValue: hValue}
		result = append(result, &lData)
		result = append(result, &hData)
	case "LL,L":
		llData := model.Level{LevelStatus: model.Status_LL_Level, LevelValue: llValue}
		lData := model.Level{LevelStatus: model.Status_L_Level, LevelValue: lValue}
		result = append(result, &llData)
		result = append(result, &lData)
	case "HH":
		hhData := model.Level{LevelStatus: model.Status_HH_Level, LevelValue: hhValue}
		result = append(result, &hhData)
	}

	return result
}
