package model

const Status_Normal = "Normal"
const Status_Error = "Error"
const Status_Start_Up = "StartUp"
const Status_Stop = "Stop"
const Status_Shut_Down = "ShutDown"
const Status_All_Open = "AllOpen"
const Status_All_Close = "AllClose"
const Status_LL_Level = "LL-Level"
const Status_L_Level = "L-Level"
const Status_M_Level = "M-Level"
const Status_H_Level = "H-Level"
const Status_HH_Level = "HH-Level"
const Status_Low_Alarm = "LowAlarm"
const Status_High_Alarm = "HighAlarm"

const Not_Status_Address = -1
const Not_Equipment_Field_Status = ""
const Not_Equipment_Field_Value = -999

const Slash = "/"
const Underline = "_"
const Colon = ":"

const Send_Log_Type = "log"
const Send_List_Type = "list"

/*(液位開關||液位計)高度*/
type Level struct {
	LevelStatus string
	LevelValue  int
}
