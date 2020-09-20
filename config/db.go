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
		Host:   GetEnv("DB_HOST", "127.0.0.1"),
		Port:   GetEnv("DB_PORT", "3306"),
		User:   GetEnv("DB_USERNAME", "root"),
		Pwd:    GetEnv("DB_PASSWORD", "root"),
		DbName: GetEnv("DB_DATABASE", "ddc"),
	}

	SlaveDbConfig = DbConf{
		Host:   GetEnv("DB_SLAVE_HOST", "127.0.0.1"),
		Port:   GetEnv("DB_SLAVE_PORT", "3306"),
		User:   GetEnv("DB_SLAVE_USERNAME", "root"),
		Pwd:    GetEnv("DB_SLAVE_PASSWORD", "root"),
		DbName: GetEnv("DB_SLAVE_DATABASE", "ddc"),
	}
)

func GetEnv(key string, value string) string {
	var val = os.Getenv(key)
	if val == "" {
		return value
	}
	return val
}
