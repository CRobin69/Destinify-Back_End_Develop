package rest

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/internal/service"
	"INTERN_BCC/pkg/helper"
	"INTERN_BCC/pkg/middleware"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:  gin.Default(),
		service: service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoint() {
	r.router.Use(r.middleware.Timeout())

	routerGroup := r.router.Group("/api/v1")

	routerGroup.GET("/health-check", healthCheck)

	routerGroup.GET("/time-out", testTimeout)

	routerGroup.GET("/login-user", r.middleware.AuthenticateUser, getLoginUser)

	routerGroup.POST("/register", r.Register)
	routerGroup.POST("/login", r.Login)

	// user.POST("/profile/upload", r.middleware.AuthenticateUser, r.UploadPhoto)

}

func (r *Rest) Serve() {
	addr := os.Getenv("APP_ADDRESS")
	port := os.Getenv("APP_PORT")

	err := r.router.Run(fmt.Sprintf("%s:%s", addr, port))
	if err != nil {
		log.Fatalf("Error while serving: %v", err)
	}
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func testTimeout(ctx *gin.Context) {
	time.Sleep(3 * time.Second)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func getLoginUser(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		helper.Error(ctx, http.StatusInternalServerError, "failed get login user", errors.New(""))
		return
	}

	helper.Success(ctx, http.StatusOK, "get login user", user.(entity.User))
}
