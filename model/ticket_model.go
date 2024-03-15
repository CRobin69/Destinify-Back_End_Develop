package model

import "github.com/google/uuid"

type TicketBuy struct {
	ID             uuid.UUID `json:"-"`
	OrderID        uuid.UUID `json:"-"`
	PlaceID        uint      `json:"placeid" binding:"required"`
	TicketPrice    int       `json:"ticket_price" binding:"required"`
	TicketQuantity int       `json:"ticket_quantity" binding:"required"`
	TotalPrice     int       `json:"-"`
	UserID         uuid.UUID `json:"userid" binding:"required"`
}

type TicketParam struct {
	ID      uuid.UUID `json:"-"`
	PlaceID uuid.UUID `json:"-"`
	UserID  uuid.UUID `json:"-"`
}
