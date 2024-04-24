package main

import (
	"log"
	"os"

	"github.com/CRobin69/Destinify-Back_End_Develop/internal/handler/rest"
	"github.com/CRobin69/Destinify-Back_End_Develop/internal/repository"
	"github.com/CRobin69/Destinify-Back_End_Develop/internal/service"
	"github.com/CRobin69/Destinify-Back_End_Develop/pkg/database/postgres"
	"github.com/CRobin69/Destinify-Back_End_Develop/pkg/middleware"
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
