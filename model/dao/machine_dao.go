package dao

import "time"

type MachineLog struct {
	Id       int       `orm:"pk;column(id)"`
	Status   string    `orm:"size(50);column(status)"`
	Datatime time.Time `orm:"type(datatime)"`
}

type MachineList struct {
	id       int
	status   string
	datatime time.Time
}
