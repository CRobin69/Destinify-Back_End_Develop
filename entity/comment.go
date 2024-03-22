package entity

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	UserID        uuid.UUID `json:"userid" gorm:"type:varchar(36);foreignKey:ID;references:users;onUpdate:CASCADE;onDelete:CASCADE"`
	TransactionID uint      `json:"transactionid" gorm:"foreignKey:ID;references:transactions;onUpdate:CASCADE;onDelete:CASCADE"`
	PlaceID       uint      `json:"placeid" gorm:"foreignKey:ID;references:places;onUpdate:CASCADE;onDelete:CASCADE"`
	StarReview    uint      `json:"star_review" gorm:"type:uint"`
	View          string    `json:"view" gorm:"type:text;not null;"`
	Feedback      string    `json:"feedback" gorm:"type:text;not null;"`
	Opinion       string    `json:"opinion" gorm:"type:text;not null;"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

