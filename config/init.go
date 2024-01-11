package config

// InitConfig TODO
func InitConfig() Config {
	cfg := new(config)
	cfg.connectDB()

	return cfg
}

type Config interface {
	Mysql() MysqlDatabase
	Port() string
}

func (di *config) Mysql() MysqlDatabase {
	return di.sqlDb
}

func (di *config) Port() string {
	return di.port
}

type config struct {
	sqlDb MysqlDatabase
	port  string
}
