package md

type KnifeGateValve struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	MultiStatus []*Status `json:"status"`
	Values      []*Value  `json:"value"`
}
