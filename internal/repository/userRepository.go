package repository

import (
	"github.com/CRobin69/Destinify-Back_End_Develop/entity"
	"github.com/CRobin69/Destinify-Back_End_Develop/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUser(param model.UserParam) (entity.User, error)
	UpdateUser(user entity.User, param model.UserParam) error
	UpdatePhoto(param model.UploadPhoto) error
	UpdatePassword(param model.UpdatePassword) error
	FindByID(id uuid.UUID) (entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(user entity.User) (entity.User, error) {
	err := ur.db.Debug().Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) GetUser(param model.UserParam) (entity.User, error) {
	user := entity.User{}
	err := ur.db.Debug().Where(&param).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) UpdateUser(user entity.User, param model.UserParam) error {
	err := ur.db.Debug().Model(&entity.User{}).Where(param.ID).Updates(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) UpdatePhoto(param model.UploadPhoto) error {
	err := ur.db.Model(&entity.User{}).Where("id = ?", param.ID).Update("profile_photo_link", param.PhotoLink).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) UpdatePassword(param model.UpdatePassword) error {
	err := ur.db.Model(&entity.User{}).Where("id = ?", param.ID).Update("password", param.NewPassword).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) FindByID(id uuid.UUID) (entity.User, error) {
	user := entity.User{}
	err := ur.db.Debug().Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
