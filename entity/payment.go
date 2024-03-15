package entity

import (
	"github.com/google/uuid"
)

type Payment struct {
	ID            uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;unique;not null;"`
	OrderID 	  uuid.UUID `json:"orderid" gorm:"type:varchar(36);foreignkey:ID;references:orders;onUpdate:CASCADE;onDelete:CASCADE"`
	Amount        int       `json:"total" gorm:"type:int;not null;"`
	PaymentDate   string    `json:"date" gorm:"type:date;not null;"`
	UserID        uuid.UUID `json:"userid" gorm:"type:varchar(36);foreignkey:ID;references:users;onUpdate:CASCADE;onDelete:CASCADE"`
	TransactionID uuid.UUID `json:"transactionid" gorm:"type:varchar(36);"`
	IsPaid        bool      `json:"isPaid" gorm:"default:false"`
	CreatedAt     string    `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     string    `json:"updatedAt" gorm:"autoUpdateTime"`
}
