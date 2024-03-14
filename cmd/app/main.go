package main

import (
	"INTERN_BCC/internal/handler/rest"
	"INTERN_BCC/internal/repository"
	"INTERN_BCC/internal/service"
	"INTERN_BCC/pkg/database/postgres"
	"INTERN_BCC/pkg/middleware"
	"os"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	env := os.Getenv("env")
	if err != nil && env == "" {
		log.Fatalf("Failed to load env, err : %v", err)
	}

	db := postgres.ConnectToDB()

	repository := repository.NewRepository(db)

	service := service.NewService(repository)

	middleware := middleware.Init(service)

	rest := rest.NewRest(service, middleware)

	postgres.Migrate(db)

	rest.MountEndpoint()
}
