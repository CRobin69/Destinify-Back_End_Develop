package entity

import (
	"time"
)

type City struct {
	ID        uint       `json:"id" gorm:"primaryKey;unique;not null"`
	Name      string     `json:"name" gorm:"type:varchar(36);not null"`
	CityImage string     `json:"cityimage" gorm:"type:text;not null"`
	Place     []Place    `json:"-"`
	Culinary  []Culinary `json:"-"`
	CreatedAt time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
}
