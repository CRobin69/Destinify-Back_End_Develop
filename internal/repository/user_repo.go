package repository

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUser(param model.UserParam) (entity.User, error)
	UpdateUser(user entity.User, param model.UserParam) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) CreateUser(user entity.User) (entity.User, error) {
	err := u.db.Debug().Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepository) GetUser(param model.UserParam) (entity.User, error) {
	user := entity.User{}
	err := u.db.Debug().Where(&param).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepository) UpdateUser(user entity.User, param model.UserParam) error {
	err := u.db.Debug().Model(&entity.User{}).Where(param).Updates(&user).Error
	if err != nil {
		return err
	}

	return nil
}
