package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load env variables
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Failed to load env, err : %v", err)
	}
}
