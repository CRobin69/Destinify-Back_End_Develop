package entity

import (
	"time"

	"github.com/google/uuid"
)

type Place struct {
	ID          uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	Name        string    `json:"name" gorm:"type:varchar(36);not null;"`
	RegionID    uuid.UUID `json:"regionid" gorm:"type:varchar(36);primary_key;foreignkey:ID;references:regions;onUpdate:CASCADE;onDelete:CASCADE"`
	Information string    `json:"information" gorm:"type:string;not null;"`
	Booking     bool      `json:"booking" default:"false"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	Ticket      []Ticket  `json:"-"`
	Guide	    []Guide   `json:"-"`
}
