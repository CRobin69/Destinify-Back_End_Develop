package rest

import (
	"encoding/json"
	"net/http"

	"INTERN_BCC/entity"
	"INTERN_BCC/model"
	"INTERN_BCC/pkg/helper"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateTransaction(ctx *gin.Context) {
	var transaction model.TransactionPost
	if err := ctx.BindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := helper.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User authentication failed"})
		return
	}

	transaction.UserID = userID.ID

	createdTransaction, err := r.service.TransactionService.CreateTransaction((transaction.UserID), transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": createdTransaction})
}

func (r *Rest) Update(ctx *gin.Context) {
	var notifPayload map[string]interface{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&notifPayload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to encode payload"})
		return
	}
	orderID, exist := notifPayload["order_id"].(string)
	if !exist {
		helper.Error(ctx, http.StatusNotFound, "order id not found", err)
		return
	}
	var data entity.Transaction
	data, err = r.service.TransactionService.Update(orderID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order ID not found"})
		return
	}

	helper.Success(ctx, http.StatusOK, "transaction updated successfully", data)
}

func (r *Rest) TransactionHistory(ctx *gin.Context) {
	var transaction model.TransactionByUserID
	user, err := helper.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User authentication failed"})
		return
	}

	transaction.ID = user.ID

	history, err := r.service.TransactionService.GetSuccessByUserID(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	helper.Success(ctx, http.StatusOK, "history showed succesfully", history)
}


