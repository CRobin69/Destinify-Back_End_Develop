package entity

import (
	"time"
)

type Place struct {
	ID            uint      `json:"id" gorm:"primary_key;unique;not null;"`
	Name          string    `json:"name" gorm:"type:varchar(100);not null;"`
	CityID        uint      `json:"cityid" gorm:"foreignkey:ID;references:cities;onUpdate:CASCADE;onDelete:CASCADE"`
	PlaceDesc     string    `json:"place_desc" gorm:"type:text;not null;"`
	PlaceAddress  string    `json:"place_address" gorm:"type:varchar(120);not null;"`
	PlaceHistory  string    `json:"place_history" gorm:"type:text;not null;"`
	PlaceFasil    string    `json:"place_fasil" gorm:"type:text;not null;"`
	PlaceActivity string    `json:"place_activity" gorm:"type:text;not null;"`
	PlaceBestTime string    `json:"place_besttime" gorm:"type:text;not null;"`
	PlaceOpen     string    `json:"place_open" gorm:"type:varchar(180);not null;"`
	PlacePrice    string    `json:"place_price" gorm:"type:varchar(200);not null;"`
	PlaceRules    string    `json:"place_rules" gorm:"type:text;not null;"`
	PlaceEvent    string    `json:"place_event" gorm:"type:text;not null;"`
	PlaceAward    string    `json:"place_award" gorm:"type:text;not null;"`
	PlaceImage    string    `json:"place_image" gorm:"type:text;not null;"`
	Price         int       `json:"price" gorm:"type:int;"`
	Ticket        []Ticket  `json:"-"`
	Guide         []Guide   `json:"-"`
	Comment       []Comment `json:"-"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
