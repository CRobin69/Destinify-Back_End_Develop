package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Status struct {
	Code      int  `json:"code"`
	IsSuccess bool `json:"is_success"`
}

func Success(ctx *gin.Context, code int, message string, data any) {
	resp := Response{
		Status: Status{
			Code:      code,
			IsSuccess: true,
		},
		Message: message,
		Data:    data,
	}
	ctx.JSON(code, resp)
}

func Error(ctx *gin.Context, code int, message string, err error) {
	resp := Response{
		Status: Status{
			Code:      code,
			IsSuccess: false,
		},
		Message: message,
		Data:    err.Error(),
	}
	ctx.JSON(code, resp)
}

func ErrorInvalidPassword() error {
	return errors.New("invalid password")
}

func ErrorEmptyEmail()error{
	return errors.New("empty email")
}

func ErrorEmptyPassword()error{
	return errors.New("empty password")
}

func ErrorEmailIsNotUnique()error{
	return errors.New("email is not unique")
}
