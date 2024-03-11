package middleware

import (
	"INTERN_BCC/internal/service"

	"github.com/gin-gonic/gin"
)

type Interface interface {
	AuthenticateUser(ctx *gin.Context)
	Timeout() gin.HandlerFunc
}

type middleware struct {
	service *service.Service
}

func Init(service *service.Service) Interface {
	return &middleware{
		service: service,
	}
}
