package model

import "github.com/google/uuid"

type CreateGuide struct {
	ID           uint   `json:"id" binding:"required"`
	PlaceID      uint   `json:"placeid" binding:"required"`
	Name         string `json:"name" binding:"required"`
	GuideDesc    string `json:"guide_desc" binding:"required"`
	GuidePrice   int    `json:"guide_price" binding:"required"`
	GuidePhoto   string `json:"guide_photo" binding:"required"`
	GuideAddress string `json:"guide_address" binding:"required"`
	GuideContact string `json:"guide_contact" binding:"required"`
}

type GuideParam struct {
	ID      uint   `json:"-"`
	PlaceID uint   `json:"-"`
	Name    string `json:"-"`
}

type GuidePatch struct {
	ID           uint      `json:"id" binding:"required"`
	Name         string    `json:"name"`
	GuideDesc    string    `json:"guide_desc"`
	GuidePrice   int       `json:"guide_price"`
	GuidePhoto   string    `json:"guide_photo"`
	GuideAddress string    `json:"guide_address"`
	GuideContact string    `json:"guide_contact"`
	Booked       bool      `json:"booked"`
}

type GuideBook struct {
	ID      uint      `json:"id" binding:"required"`
	PlaceID uint      `json:"placeid" binding:"required"`
	Booked  bool      `json:"booked"`
	UserID  uuid.UUID `json:"-"`
}
