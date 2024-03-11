package entity

import (
	"time"

	"github.com/google/uuid"
)

type Region struct {
	ID        uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	Name      string    `json:"name" gorm:"type:varchar(36);not null;"`
	Data      string    `json:"data" gorm:"type:string;not null;"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	Place	 []Place    `json:"place"`
}
