package dao

import "time"

type FloatLevelSwitchLog struct {
	Id       int       `orm:"pk;column(id)"`
	Level    string    `orm:"size(50);column(level)"`
	Status   string    `orm:"size(50);column(status)"`
	Datatime time.Time `orm:"type(datatime)"`
}

type FloatLevelSwitchList struct {
	id       int
	level    string
	status   string
	datatime time.Time
}
