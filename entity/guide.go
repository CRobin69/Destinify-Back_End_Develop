package entity

import (
	"time"

	"github.com/google/uuid"
)

type Guide struct {
	ID        uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	UserID    uuid.UUID `json:"userid" gorm:"type:varchar(36);primary_key;foreignkey:ID;references:users;onUpdate:CASCADE;onDelete:CASCADE"`
	Name      string    `json:"name" gorm:"type:varchar(36);not null;"`
	Price     int       `json:"price" gorm:"type:int;not null;"`
	PlaceID   uuid.UUID `json:"placeid" gorm:"type:varchar(36);primary_key;foreignkey:ID;references:places;onUpdate:CASCADE;onDelete:CASCADE"`
	PaymentID uuid.UUID `json:"paymentid" gorm:"type:varchar(36);primary_key;foreignkey:ID;references:payments;onUpdate:CASCADE;onDelete:CASCADE"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
