package md

type LiquidLevelGauge struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	MultiStatus []*Status `json:"status"`
	Values      []*Value  `json:"value"`
	Level_mode  string    `json:"level_mode"`
}
