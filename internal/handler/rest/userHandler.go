package rest

import (
	"INTERN_BCC/model"
	"INTERN_BCC/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	param := model.UploadPhoto{}
	photo, err := ctx.FormFile("photo")
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	user, err := helper.GetLoginUser(ctx)
	
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User authentication failed"})
		return
	}
	param.ID = user.ID
	photoLink, uploadErr := r.service.UserService.UploadPhoto(model.UploadPhoto{Photo: photo})
	if uploadErr != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to upload photo", uploadErr)
		return
	}

	helper.Success(ctx, http.StatusOK, "success upload photo", photoLink)
}

func (r *Rest) UpdateUser(ctx *gin.Context) {
	param := model.UpdateUser{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}
	user, err := helper.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User authentication failed"})
		return
	}
	param.ID = user.ID
	_, err = r.service.UserService.UpdateUser(param.ID)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to update user", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success update user", nil)
}

func (r *Rest) UpdatePassword(ctx *gin.Context) {
	param := model.UpdatePassword{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}
	user, err := helper.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User authentication failed"})
		return
	}
	param.ID = user.ID
	_, err = r.service.UserService.UpdatePassword(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to update password", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success update password", nil)
}

func (r *Rest) GetUserByID(ctx *gin.Context) {
	param := model.UserParam{}

	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	realUserID, ok := userID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	param.ID = realUserID

	user, err := r.service.UserService.GetUser(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to get user", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success get user", user)
}
