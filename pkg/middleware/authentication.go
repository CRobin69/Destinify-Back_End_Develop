package middleware

import (
	"github.com/CRobin69/Destinify-Back_End_Develop/model"
	"github.com/CRobin69/Destinify-Back_End_Develop/pkg/helper"
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
	userID, err := helper.ValidateToken(token)
	if err != nil {
		helper.Error(ctx, http.StatusUnauthorized, "failed validate token", err)
		ctx.Abort()
		return
	}

	ctx.Set("user_id", userID)
	user, err := m.service.UserService.GetUser(model.UserParam{
		ID: userID,
	})
	if err != nil {
		helper.Error(ctx, http.StatusUnauthorized, "failed get user", err)
		ctx.Abort()
	}

	ctx.Set("user", user)

	ctx.Next()
}
