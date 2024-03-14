package model

import "github.com/google/uuid"


type TicketCreate struct {
	ID          	uuid.UUID `json:"-"`
	PlaceID     	uint 	  `json:"placeid" binding:"required"`
	TicketPrice 	int   	  `json:"ticket_price" binding:"required"`
	TicketDate  	string    `json:"ticket_date" binding:"required"`
	TicketQuantity 	int    	  `json:"ticket_quantity" binding:"required"`
	TotalPrice  	int   	  `json:"total_price"`
	UserID      	uuid.UUID `json:"userid" binding:"required"`
}

type TicketParam struct {
	ID    	uuid.UUID `json:"-"`
	PlaceID uuid.UUID `json:"placeid" binding:"required"`
	UserID  uuid.UUID `json:"userid" binding:"required"`
}