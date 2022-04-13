package md

type Param struct {
	CronRealTimeSpec string   `json:"cron_real_time_spec"`
	CronHistorySpec  string   `json:"cron_history_spec"`
	ModbusIp         string   `json:"modbus_ip"`
	Redis            Redis    `json:"redis"`
	Postgres         Postgres `json:"postgres"`
}

type Redis struct {
	Ip            string   `json:"ip"`
	Pwd           string   `json:"pwd"`
	Db            int      `json:"db"`
	EquipmentKeys []string `json:"equipment_key"`
}

type Postgres struct {
	DataSource string `json:"data_source"`
}
