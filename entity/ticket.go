package entity

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID          uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;unique;not null;"`
	PlaceID     uint      `json:"placeid" gorm:"foreignkey:ID;references:places;onUpdate:CASCADE;onDelete:CASCADE"`
	UserID      uuid.UUID `json:"userid" gorm:"type:varchar(36);foreignkey:ID;references:users;onUpdate:CASCADE;onDelete:CASCADE"`
	OrderID     uuid.UUID `json:"orderid" gorm:"type:varchar(36);foreignkey:ID;references:orders;onUpdate:CASCADE;onDelete:CASCADE"`
	TicketPrice int       `json:"ticket_price" gorm:"type:int;not null;"`
	TicketDate  time.Time `json:"ticket_date" gorm:"autoCreateTime;not null;"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
