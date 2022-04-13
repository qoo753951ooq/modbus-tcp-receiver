package db

import (
	"github.com/astaxie/beego/orm"

	_ "github.com/lib/pq"
)

func InitPostgresDB(dataSoure string) {

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", dataSoure)
	orm.SetMaxIdleConns("default", 20)
	orm.SetMaxOpenConns("default", 20)

	//orm.RegisterModel()
	//orm.RunSyncdb("default", false, true)
	orm.Debug = false
}
