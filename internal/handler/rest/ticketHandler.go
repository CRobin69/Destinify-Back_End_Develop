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

func (r *Rest) BuyTicket(ctx *gin.Context) {
	var ticketBuy model.TicketBuy
	if err := ctx.ShouldBindJSON(&ticketBuy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract the userID from the context
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	// userID is an interface{}, so you'll need to cast it to its original type
	realUserID := userID.(uuid.UUID)

	// Update the UserID field in the ticketBuy model
	ticketBuy.UserID = realUserID

	// Call the BuyTicket function in the TicketService
	tickets, err := r.service.TicketService.BuyTicket(ticketBuy)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to buy tickets", "details": err.Error()})
		return
	}

	// Create a slice to hold the ticket IDs
	var ticketIDs []uuid.UUID
	for _, ticket := range tickets {
		ticketIDs = append(ticketIDs,ticket.ID)
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Tickets created successfully", "TicketIDs": ticketIDs})
}
