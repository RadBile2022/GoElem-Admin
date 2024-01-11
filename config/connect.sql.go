package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"sync"
)

func (di *config) connectDB() {
	dotenv := ".env"

	err := godotenv.Load(dotenv)
	if fileExist(dotenv) && err == nil {
		hostMysql, there = os.LookupEnv("HOST_MYSQL")
		portMysql, there = os.LookupEnv("PORT_MYSQL")
		dbMysql, there = os.LookupEnv("DB_MYSQL")
		userMysql, there = os.LookupEnv("USER_MYSQL")
		passMysql, there = os.LookupEnv("PASS_MYSQL")
		portApp, there = os.LookupEnv("PORT")
	} else {
		fmt.Println("Error -:- file .env doesn't found !")
		return
	}

	env := valEnv{
		Host:     hostMysql,
		Port:     portMysql,
		DB:       dbMysql,
		Username: userMysql,
		Password: passMysql,
	}

	var loadOnce sync.Once
	loadOnce.Do(func() {
		di.sqlDb = InitMysql(env)
		di.port = portApp
	})
}

func fileExist(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

var (
	hostMysql string
	portMysql string
	dbMysql   string
	userMysql string
	passMysql string
	portApp   string
	there     bool
)
