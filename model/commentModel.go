package model

import "github.com/google/uuid"

type CommentCreate struct {
	ID            uint      `json:"-"`
	UserID        uuid.UUID `json:"-"`
	StarReview    uint      `json:"star_review" binding:"required"`
	View          string    `json:"view" binding:"required"`
	Feedback      string    `json:"feedback" binding:"required"`
	Opinion       string    `json:"opinion" binding:"required"`
	TransactionID uint      `json:"-"`
	PlaceID       uint      `json:"-"`
}

type CommentParam struct {
	ID      uint      `json:"-"`
	UserID  uuid.UUID `json:"-"`
	PlaceID uint      `json:""`
}
