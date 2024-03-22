package config

import (
	"os"
)

func LoadDatabaseConfig() string {
	dsn := os.Getenv("DB_URL")
	return dsn
}
