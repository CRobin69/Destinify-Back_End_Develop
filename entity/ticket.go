package entity

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID        uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	PlaceID   uuid.UUID `json:"placeid" gorm:"type:varchar(36);primary_key;foreignkey:ID;references:places;onUpdate:CASCADE;onDelete:CASCADE"`
	Price     int       `json:"price" gorm:"type:int;not null;"`
	Date      string    `json:"date" gorm:"type:date;not null;"`
	Total     int       `json:"total" gorm:"type:int;not null;"`
	UserID    uuid.UUID `json:"userid" gorm:"type:varchar(36);primary_key;foreignkey:ID;references:users;onUpdate:CASCADE;onDelete:CASCADE"`
	PaymentID uuid.UUID `json:"paymentid" gorm:"type:varchar(36);primary_key;foreignkey:ID;references:payments;onUpdate:CASCADE;onDelete:CASCADE"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
