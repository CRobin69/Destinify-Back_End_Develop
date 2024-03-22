package entity

import (
	"time"


)

type Culinary struct {
	ID                 uint  	  `json:"id" gorm:"primary_key"`
	Name               string     `json:"name" gorm:"type:varchar(36);not null"`
	CityID             uint       `json:"cityid" gorm:"foreignkey:ID;references:cities;onUpdate:CASCADE;onDelete:CASCADE"`
	CulinaryDesc       string     `json:"culinary_desc" gorm:"type:text;not null;"`
	CulinaryAddress    string     `json:"culinary_address" gorm:"type:varchar(120);not null;"`
	CulinaryPriceRange string     `json:"culinary_price_range" gorm:"type:varchar(50);not null;"`
	CulinaryOpen       string     `json:"culinary_open" gorm:"type:varchar(10);not null;"`
	CulinaryClose      string     `json:"culinary_close" gorm:"type:varchar(10);not null;"`
	CulinaryImage      string     `json:"culinary_image" gorm:"type:text;not null;"`
	CulinaryAward      string     `json:"culinary_award" gorm:"type:text;not null;"`
	CreatedAt          time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt          time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
}
