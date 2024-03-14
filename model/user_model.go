package model

import (
	"github.com/google/uuid"
)

type UserRegister struct {
	ID              uuid.UUID `json:"-"`
	HP    			string    `json:"hp" binding:"required"`
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

// type UserUploadPhoto struct {
// 	Photo *multipart.FileHeader `form:"photo"`
// }
