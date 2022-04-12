package dao

import (
	"time"
)

type KnifeGateValveLog struct {
	Id       int       `orm:"pk;column(id)"`
	Status   string    `orm:"size(50);column(status)"`
	Opening  float64   `orm:"null;column(opening)"`
	Datatime time.Time `orm:"type(datatime)"`
}

type KnifeGateValveList struct {
	id       int
	status   string
	opening  float64
	datatime time.Time
}
