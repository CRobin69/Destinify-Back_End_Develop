package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID            uint      `json:"id" gorm:"primary_key;unique;not null;"`
	UserID        uuid.UUID `json:"userid" gorm:"type:varchar(36);foreignkey:ID;references:users;onUpdate:CASCADE;onDelete:CASCADE"`
	OrderID       uuid.UUID `json:"orderid" gorm:"type:varchar(36);foreignkey:ID;references:orders;onUpdate:CASCADE;onDelete:CASCADE"`
	TransactionID string    `json:"transactionid" gorm:"type:varchar(100);unique;"`
	PlaceID       uint      `json:"placeid" gorm:"foreignkey:ID;references:places;onUpdate:CASCADE;onDelete:CASCADE"`
	Amount        int       `json:"total" gorm:"type:int;not null;"`
	Method        string    `json:"method" gorm:"type:varchar(100);not null;"`
	VANumber      string    `json:"va_number" gorm:"type:varchar(100);"`
	Status        string    `json:"status" gorm:"type:varchar(100);not null;"`
	Place         Place     `json:"place"`
	Comment	      Comment 	`json:"-"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
