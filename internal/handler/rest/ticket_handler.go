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

func (r *Rest) BuyTicket(c *gin.Context) {
	var ticketBuy model.TicketBuy
	if err := c.BindJSON(&ticketBuy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tickets, err := r.service.TicketService.BuyTicket(ticketBuy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to buy tickets", "details": err.Error()})
		return
	}

	for _, ticket := range tickets {
		c.JSON(http.StatusOK, gin.H{"message": "ticket created successfully", "TicketID": ticket.ID})
	}
}