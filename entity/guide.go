package entity

import (
	"time"

	"github.com/google/uuid"
)

type Guide struct {
	ID           uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	Name         string    `json:"name" gorm:"type:varchar(36);not null;"`
	UserID       uuid.UUID `json:"userid" gorm:"type:varchar(36);foreignkey:ID;references:users;onUpdate:CASCADE;onDelete:CASCADE"`
	PlaceID      uint      `json:"placeid" gorm:"foreignkey:ID;references:places;onUpdate:CASCADE;onDelete:CASCADE"`
	PaymentID    uuid.UUID `json:"paymentid" gorm:"type:varchar(36);foreignkey:ID;references:payments;onUpdate:CASCADE;onDelete:SET NULL"`
	GuideDesc    string    `json:"guide_desc" gorm:"type:varchar(1200);not null;"`
	GuidePrice   int       `json:"guide_price" gorm:"type:int;not null;"`
	GuidePhoto   string    `json:"guide_photo" gorm:"type:varchar(500);not null;"`
	GuideAddress string    `json:"guide_address" gorm:"type:varchar(120);not null;"`
	GuideContact string    `json:"guide_contact" gorm:"type:varchar(15);not null;"`
	Price        int       `json:"price" gorm:"type:int;not null;"`
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
