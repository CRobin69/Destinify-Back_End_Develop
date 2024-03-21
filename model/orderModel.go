package model

import "github.com/google/uuid"

type CreateOrder struct {
	ID     uuid.UUID `json:"-"`
	UserID uuid.UUID `json:"-"`
	
}
