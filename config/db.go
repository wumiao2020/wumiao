package config

import (
	"os"
)

const DriverName = "mysql"

type DbConf struct {
	Host   string
	Port   string
	User   string
	Pwd    string
	DbName string
}

var (
	MasterDbConfig = DbConf{
		Host:   getEnv("DB_HOST", "127.0.0.1"),
		Port:   getEnv("DB_PORT", "3306"),
		User:   getEnv("DB_USERNAME", "root"),
		Pwd:    getEnv("DB_PASSWORD", "root"),
		DbName: getEnv("DB_DATABASE", "ddc"),
	}

	SlaveDbConfig = DbConf{
		Host:   getEnv("DB_SLAVE_HOST", "127.0.0.1"),
		Port:   getEnv("DB_SLAVE_PORT", "3306"),
		User:   getEnv("DB_SLAVE_USERNAME", "root"),
		Pwd:    getEnv("DB_SLAVE_PASSWORD", "root"),
		DbName: getEnv("DB_SLAVE_DATABASE", "ddc"),
	}
)

func getEnv(key string, value string) string {
	var val = os.Getenv(key)
	if val == "" {
		return value
	}
	return val
}
