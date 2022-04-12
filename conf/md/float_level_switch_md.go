package md

type FloatLevelSwitch struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	MultiStatus []*Status `json:"status"`
	Level_mode  string    `json:"level_mode"`
}
