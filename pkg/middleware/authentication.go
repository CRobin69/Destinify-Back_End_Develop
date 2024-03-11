package middleware

import (
	"INTERN_BCC/model"
	"INTERN_BCC/pkg/helper"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	
)

func (m *middleware) AuthenticateUser(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	if bearer == "" {
		helper.Error(ctx, http.StatusUnauthorized, "empty token", errors.New(""))
		ctx.Abort()
	}

	token := strings.Split(bearer, " ")[1]
	userId, err := helper.ValidateToken(token)
	if err != nil {
		helper.Error(ctx, http.StatusUnauthorized, "failed validate token", err)
		ctx.Abort()
	}

	user, err := m.service.UserService.GetUser(model.UserParam{
		ID: userId,
	})
	if err != nil {
		helper.Error(ctx, http.StatusUnauthorized, "failed get user", err)
		ctx.Abort()
	}

	ctx.Set("user", user)

	ctx.Next()
}
