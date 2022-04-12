package md

type Pump struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	MultiStatus []*Status `json:"status"`
}
