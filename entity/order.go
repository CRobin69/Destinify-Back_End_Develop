package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID        uuid.UUID   `json:"id" gorm:"type:varchar(36);primaryKey"`
	UserID    uuid.UUID   `json:"userid" gorm:"type:varchar(36);foreignkey:ID;references:users;onUpdate:CASCADE;onDelete:CASCADE"`
	GuideID   uuid.UUID   `json:"guide_id" gorm:"type:varchar(36);foreignkey:ID;references:guides;onUpdate:CASCADE;onDelete:CASCADE"`
	CreatedAt time.Time   `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time   `json:"updatedAt" gorm:"autoUpdateTime"`
	TicketID  []uuid.UUID `json:"ticket_ids" gorm:"type:varchar(36);foreignkey:ID;references:tickets;onUpdate:CASCADE;onDelete:CASCADE"`
	Payment   []Payment   `json:"-"`
}

