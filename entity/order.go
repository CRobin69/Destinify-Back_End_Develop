package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID   `json:"id" gorm:"type:varchar(36);primaryKey"`
	UserID      uuid.UUID   `json:"userid" gorm:"type:varchar(36);foreignkey:ID;references:users;onUpdate:CASCADE;onDelete:CASCADE"`
	GuideID     uint        `json:"guide_id" gorm:"foreignkey:ID;references:guides;onUpdate:CASCADE;onDelete:SET NULL"`
	TotalPrice  int         `json:"total_price" gorm:"type:int;"`
	Tickets     []uuid.UUID `json:"tickets" gorm:"type:uuid[]"`
	Transaction Transaction `json:"-"`
	Guide       Guide       `json:"-"`
	CreatedAt   time.Time   `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `json:"updatedAt" gorm:"autoUpdateTime"`
}
