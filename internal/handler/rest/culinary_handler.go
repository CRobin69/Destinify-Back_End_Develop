package rest

import (
	"INTERN_BCC/model"
	"INTERN_BCC/pkg/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateCulinary(ctx *gin.Context) {
	param := model.CulinaryCreate{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = r.service.CulinaryService.CreateData(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to create culinary", err)
		return
	}

	helper.Success(ctx, http.StatusCreated, "success create culinary", nil)
}

func (r *Rest) GetCulinaryByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "invalid id", err)
		return
	}
	param := model.CulinaryParam{ID: uint(idUint)}
	culinary, err := r.service.CulinaryService.GetCulinaryByID(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to get culinary", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success get culinary", culinary)
}

func (r *Rest) GetAllCulinary(ctx *gin.Context) {
	param := model.CulinaryParam{}
	culinary, err := r.service.CulinaryService.GetAllCulinary(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to get all culinary", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success get all culinary", culinary)
}

// func (r *Rest) SearchCulinary(ctx *gin.Context) {
// 	param := model.SearchCulinary{}
// 	culinary, err := r.service.CulinaryService.SearchCulinary(param)
// 	if err != nil {
// 		helper.Error(ctx, http.StatusInternalServerError, "failed to search culinary", err)
// 		return
// 	}
//
// 	helper.Success(ctx, http.StatusOK, "success search culinary", culinary)
// }
