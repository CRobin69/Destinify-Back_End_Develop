package rest

import (
	"INTERN_BCC/model"
	"INTERN_BCC/pkg/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetCommentByPlaceID(ctx *gin.Context) {
	id := ctx.Param("placeid")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "invalid id", err)
		return
	}

	param := model.CommentParam{PlaceID: uint(idUint)}
	comments, err := r.service.CommentService.FindCommentByPlaceID(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to get comment by city id", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success get comments by place id", comments)
}

func (r *Rest) GetCommentByUserID(ctx *gin.Context) {
	var param model.CommentParam
	user, err := helper.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User authentication failed"})
		return
	}
	param.UserID = user.ID
	comments, err := r.service.CommentService.FindCommentByUserID(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to get comment by city id", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success get comments by user id", comments)
}

func (r *Rest) UpdateComment(ctx *gin.Context) {
	var param model.CommentCreate
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "invalid id", err)
		return
	}
	param.ID = uint(idUint)
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := helper.GetLoginUser(ctx)
	if err != nil {
		helper.Error(ctx, http.StatusUnauthorized,  "User authentication failed", err)
		return
	}
	param.UserID = user.ID
	err = r.service.CommentService.UpdateComment(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to update guide", err)
		return
	}
	helper.Success(ctx, http.StatusOK, "success update comments", nil)
}


