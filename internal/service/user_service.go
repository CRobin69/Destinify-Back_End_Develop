package service

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/internal/repository"
	"INTERN_BCC/model"
	"INTERN_BCC/pkg/helper"
	"github.com/google/uuid"
)

type IUserService interface {
	Register(param model.UserRegister) error
	GetUser(param model.UserParam) (entity.User, error)
	Login(param model.UserLogin) (model.UserLoginResponse, error)
}

type UserService struct {
	ur      repository.IUserRepository
	
	// supabase supabase.Interface
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	// supabase supabase.Interface
	return &UserService{
		ur: userRepository,
	
	}
	// supabase: supabase,

}

func (u *UserService) Register(param model.UserRegister) error {
	hashPassword, err := helper.HashPassword(param.Password)
	if err != nil {
		return err
	}
	param.ID = uuid.New()
	param.Password = hashPassword

	user := entity.User{
		ID:       param.ID,
		HP: 	  param.HP,
		Name:     param.Name,
		Email:    param.Email,
		Password: param.Password,
	}

	_, err = u.ur.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) Login(param model.UserLogin) (model.UserLoginResponse, error) {
	result := model.UserLoginResponse{}

	user, err := u.ur.GetUser(model.UserParam{
		Email: param.Email,
	})
	if err != nil {
		return result, err
	}

	err = helper.ComparePassword(user.Password, param.Password)
	if err != nil {
		return result, err
	}

	token, err := helper.CreateJWTToken(user.ID)
	if err != nil {
		return result, err
	}

	result.Token = token

	return result, nil
}

func (u *UserService) GetUser(param model.UserParam) (entity.User, error) {
	return u.ur.GetUser(param)
}

// func (u *UserService) UploadPhoto(ctx *gin.Context, param model.UserUploadPhoto) error {
// 	user, err := u.jwtAuth.GetLoginUser(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	if user.PhotoLink != "" {
// 		err := u.supabase.Delete(user.PhotoLink)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	link, err := u.supabase.Upload(param.Photo)
// 	if err != nil {
// 		return err
// 	}

// 	err = u.ur.UpdateUser(entity.User{
// 		PhotoLink: link,
// 	}, model.UserParam{
// 		ID: user.ID,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
