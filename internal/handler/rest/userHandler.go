package rest

import (
	"INTERN_BCC/model"
	"INTERN_BCC/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) Register(ctx *gin.Context) {
	RegisterParam := model.UserRegister{}

	err := ctx.ShouldBindJSON(&RegisterParam)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	if RegisterParam.Password != RegisterParam.ConfirmPassword {
		helper.Error(ctx, http.StatusBadRequest, "password and confirm password not match", nil)
		return
	}

	err = r.service.UserService.Register(RegisterParam)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to register new user", err)
		return
	}

	helper.Success(ctx, http.StatusCreated, "success register new user", nil)
}

func (r *Rest) Login(ctx *gin.Context) {
	param := model.UserLogin{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	token, err := r.service.UserService.Login(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to login", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success login to system", token)
}

func (r *Rest) UploadPhoto(ctx *gin.Context) {
	photo, err := ctx.FormFile("photo")
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	photoLink, uploadErr := r.service.UserService.UploadPhoto(ctx, model.UploadPhoto{Photo: photo})
	if uploadErr != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to upload photo", uploadErr)
		return
	}

	helper.Success(ctx, http.StatusOK, "success upload photo", photoLink)
}