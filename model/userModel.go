package model

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type UserRegister struct {
	ID              uuid.UUID `json:"-"`
	HP              string    `json:"hp" binding:"required"`
	Name            string    `json:"name" binding:"required"`
	Email           string    `json:"email" binding:"required,email"`
	Password        string    `json:"password" binding:"required,min=8"`
	ConfirmPassword string    `json:"confirm_password" binding:"required,eqfield=Password"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserParam struct {
	ID       uuid.UUID `json:"-"`
	Email    string    `json:"-"`
	Password string    `json:"-"`
}

type UploadPhoto struct {
	ID        uuid.UUID             `json:"-"`
	PhotoLink string                `json:"-"`
	Photo     *multipart.FileHeader `form:"photo"`
}

type UpdateUser struct {
	ID              uuid.UUID `json:"-"`
	HP              string    `json:"hp"`
	Name            string    `json:"name"`
	Email           string    `json:"-"`
}

type UpdatePassword struct {
	ID              uuid.UUID `json:"-"`
	OldPassword     string    `json:"old_password" binding:"required"`
	NewPassword     string    `json:"new_password" binding:"required,min=8"`
	ConfirmPassword string    `json:"confirm_password" binding:"required,eqfield=NewPassword"`
}
