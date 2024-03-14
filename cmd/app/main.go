package main

import (
	"INTERN_BCC/internal/handler/rest"
	"INTERN_BCC/internal/repository"
	"INTERN_BCC/internal/service"
	"INTERN_BCC/pkg/config"
	"INTERN_BCC/pkg/database/postgres"
	"INTERN_BCC/pkg/middleware"
)

func main() {
	config.LoadEnv()
	db := postgres.ConnectToDB()

	repository := repository.NewRepository(db)

	service := service.NewService(repository)

	middleware := middleware.Init(service)

	rest := rest.NewRest(service, middleware)

	postgres.Migrate(db)

	rest.MountEndpoint()
}
