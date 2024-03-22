package rest

import (
	"INTERN_BCC/model"
	"INTERN_BCC/pkg/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateGuide(ctx *gin.Context) {
	param := model.CreateGuide{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}
	err = r.service.GuideService.CreateGuide(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to create guide", err)
		return
	}

	helper.Success(ctx, http.StatusCreated, "success create guide", nil)
}

func (r *Rest) PatchGuide(ctx *gin.Context) {
	param := model.GuidePatch{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}
	err = r.service.GuideService.PatchGuide(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to update guide", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success update guide", nil)
}

func (r *Rest) GetAllGuide(ctx *gin.Context) {
	param := model.GuideParam{}
	guides, err := r.service.GuideService.GetAllGuide(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to get all guide", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success get all guide", guides)
}

func (r *Rest) GetGuideByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "invalid id", err)
		return
	}
	param := model.GuideParam{ID: uint(idUint)}
	guide, err := r.service.GuideService.GetGuideByID(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to get guide", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success get guide", guide)
}


