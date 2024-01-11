package config

import (
	"database/sql"
	"fmt"
	"github.com/CezarGarrido/sqllogs"
	"github.com/labstack/gommon/log"
	"time"
)

// InitMysql TODO a2 : mysqlConfig
func InitMysql(val valEnv) MysqlDatabase {
	mysql := new(mysqlDatabase)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		val.Username, val.Password, val.Host, val.Port, val.DB)

	// Open up our database connection. If there is an error opening the connection, handle it
	sqllogs.SetDebug(true)
	conn, err := sql.Open("sqllog:mysql", connection)
	if err != nil {
		fmt.Println(err, "Error :-: Connection DB is Failed")
	}

	err = conn.Ping()
	if err != nil {
		fmt.Println(err, "Error :-: Ping Connection DB is Failed")
	}

	conn.SetMaxOpenConns(30)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(3 * time.Minute)

	mysql.sqlDB = conn

	log.Info("Mysql Successfully Connected . . .")

	return mysql
}

type MysqlDatabase interface {
	GetConnection() *sql.DB
}

func (di *mysqlDatabase) GetConnection() *sql.DB {
	return di.sqlDB
}

type mysqlDatabase struct {
	sqlDB *sql.DB
}

type valEnv struct {
	Host     string
	Port     string
	DB       string
	Username string
	Password string
}
