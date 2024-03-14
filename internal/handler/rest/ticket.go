package rest

import (
	"INTERN_BCC/model"
	"INTERN_BCC/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) GetTicketByID(ctx *gin.Context) {
	id := ctx.Param("id")
	param := model.TicketParam{ID: uuid.MustParse(id)}
	ticket, err := r.service.TicketService.GetTicketByID(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to get ticket", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success get ticket", ticket)
}

func (r *Rest) CreateTicket(ctx *gin.Context) {
	param := model.TicketCreate{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		helper.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = r.service.TicketService.CreateTicket(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to create ticket", err)
		return
	}

	helper.Success(ctx, http.StatusCreated, "success create ticket", nil)
}
