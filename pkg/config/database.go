package config

import (
	"os"
)

func LoadDatabaseConfig() string {
	dsn := os.Getenv("DB_URL")
	return dsn
	//  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
	// os.Getenv("DB_HOST"),
	// os.Getenv("DB_USER"),
	// os.Getenv("DB_PASSWORD"),
	// os.Getenv("DB_NAME"),
	// os.Getenv("DB_PORT"),
	// )
}
