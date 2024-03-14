package model

import "github.com/google/uuid"

type CulinaryCreate struct {
	ID                 uuid.UUID `json:"-"`
	Name               string    `json:"name" binding:"required"`
	CityID             uint      `json:"cityid" binding:"required"`
	CulinaryDesc       string    `json:"culinary_desc" binding:"required"`
	CulinaryAddress    string    `json:"culinary_address" binding:"required"`
	CulinaryPriceRange string    `json:"culinary_price_range" binding:"required"`
	CulinaryOpen       string    `json:"culinary_open" binding:"required"`
	CulinaryClose      string    `json:"culinary_close" binding:"required"`
	CulinaryImage      string    `json:"culinary_image" binding:"required"`
	CulinaryAward      string    `json:"culinary_award" binding:"required"`
}

type CulinaryParam struct {
	ID     uuid.UUID `json:"-"`
	Name   string    `json:"name"`
	CityID uint      `json:"cityid"`
}

type SearchCulinary struct {
	Name string `json:"name"`
}