package rest

import (
	"INTERN_BCC/model"
	"INTERN_BCC/pkg/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateCity(ctx *gin.Context) {
	param := model.CityCreate{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = r.service.CityService.CreateCity(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to create city", err)
		return
	}

	helper.Success(ctx, http.StatusCreated, "success create city", nil)
}

func (r *Rest) GetCity(ctx *gin.Context) {
	id := ctx.Param("id")
	cityID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "invalid city ID", err)
		return
	}

	cities, err := r.service.CityService.GetCity(model.CityParam{ID: uint(cityID)})
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to get city", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success get city", cities)
}

func (r *Rest) GetAllCity(ctx *gin.Context) {
	cities, err := r.service.CityService.GetAllCity(model.CityParam{})
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to get all city", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success get all city", cities)
}

func (r *Rest) SearchCity(ctx *gin.Context) {
	param := model.SearchCity{
		Name: ctx.Query("name"),
	}
	cities, err := r.service.CityService.SearchCity(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to search city", err)
		return
	}
	if len(cities) == 0 {
		helper.Success(ctx, http.StatusOK, "no city found", nil)
		return
	}

	helper.Success(ctx, http.StatusOK, "success search city", cities)
}
