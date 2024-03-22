package model

import (
	"github.com/google/uuid"
)

type TransactionByID struct {
	ID       uint   `json:"id"`
	Amount   int    `json:"total"`
	Method   string `json:"method"`
	VANumber string `json:"va_number"`
	OrderID  string `json:"orderid"`
	Status   string `json:"status"`
}

type TransactionByUserID struct {
	ID uuid.UUID `json:"-"`
}

type TransactionPost struct {
	UserID  uuid.UUID `json:"-"`
	Amount  int       `json:"-"`
	Method  string    `json:"method" binding:"required"`
	OrderID uuid.UUID `json:"-"`
}
