package model

import (
	"time"

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

type TransactionByUser struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Pict      string    `json:"pict"`
	Title     string    `json:"title"`
}

type TransactionPost struct {
	UserID  uuid.UUID `json:"-"`
	Amount  int       `json:"amount"`
	Method  string    `json:"method" binding:"required"`
	OrderID uuid.UUID `json:"orderid"`
}
