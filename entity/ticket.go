package entity

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID             uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	PlaceID        uint      `json:"placeid" gorm:"foreignkey:ID;references:places;onUpdate:CASCADE;onDelete:CASCADE"`
	UserID         uuid.UUID `json:"userid" gorm:"type:varchar(36);primary_key;foreignkey:ID;references:users;onUpdate:CASCADE;onDelete:CASCADE"`
	PaymentID      uuid.UUID `json:"paymentid" gorm:"type:varchar(36);primary_key;foreignkey:ID;references:payments;onUpdate:CASCADE;onDelete:SET NULL"`
	TicketPrice    int       `json:"ticket_price" gorm:"type:int;not null;"`
	TicketQuantity int       `json:"ticket_quantity" gorm:"type:int;not null;"`
	TicketDate     string    `json:"ticket_date" gorm:"type:date;not null;"`
	TotalPrice     int       `json:"total_price" gorm:"type:int;not null;"`
	CreatedAt      time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
