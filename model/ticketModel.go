package model

import "github.com/google/uuid"

type TicketBuy struct {
	ID             uuid.UUID `json:"-"`
	OrderID        uuid.UUID `json:"-"`
	UserID         uuid.UUID `json:"-"`
	GuideID 	   uint      `json:"guideid"`
	PlaceID        uint      `json:"placeid" binding:"required"`
	TicketPrice    int       `json:"ticket_price"`
	GuidePrice     int       `json:"guide_price"`
	TicketQuantity int       `json:"ticket_quantity" binding:"required"`
}

type TicketParam struct {
	ID      uuid.UUID `json:"-"`
	PlaceID uuid.UUID `json:"-"`
	UserID  uuid.UUID `json:"-"`
}
