package service

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/internal/repository"
	"INTERN_BCC/model"
	"INTERN_BCC/pkg/helper"
	"INTERN_BCC/pkg/supabase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IUserService interface {
	Register(param model.UserRegister) error
	GetUser(param model.UserParam) (entity.User, error)
	Login(param model.UserLogin) (model.UserLoginResponse, error)
	UploadPhoto(ctx *gin.Context, param model.UploadPhoto) (string, error)
	UpdateUser(ctx *gin.Context, param model.UpdateUser) (string, error)
	UpdatePassword(ctx *gin.Context, param model.UpdatePassword) (string, error)
}

type UserService struct {
	ur repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{
		ur: userRepository,
	}

}

func (us *UserService) Register(param model.UserRegister) error {
	hashPassword, err := helper.HashPassword(param.Password)
	if err != nil {
		return err
	}
	param.ID = uuid.New()
	param.Password = hashPassword

	user := entity.User{
		ID:       param.ID,
		HP:       param.HP,
		Name:     param.Name,
		Email:    param.Email,
		Password: param.Password,
	}

	_, err = us.ur.CreateUser(user)
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

func (u *UserService) UploadPhoto(ctx *gin.Context, param model.UploadPhoto) (string, error) {
	user, err := helper.GetLoginUser(ctx)
	if err != nil {
		return "", err
	}

	supabaseStorage := supabase.NewSupabaseStorage()

	if user.PhotoLink != "" {
		err := supabaseStorage.Delete(user.PhotoLink)
		if err != nil {
			return "", err
		}
	}

	link, err := supabaseStorage.Upload(param.Photo)
	if err != nil {
		return "", err
	}

	err = u.ur.UpdateUser(entity.User{
		PhotoLink: link,
	}, model.UserParam{
		ID: user.ID,
	})
	if err != nil {
		return "", err
	}

	return link, nil
}

func (u *UserService) UpdateUser(ctx *gin.Context, param model.UpdateUser) (string, error) {
	user, err := helper.GetLoginUser(ctx)
	if err != nil {
		return "", err
	}

	if param.HP != "" {
		user.HP = param.HP
	}
	if param.Name != "" {
		user.Name = param.Name
	}

	err = u.ur.UpdateUser(user, model.UserParam{
		ID: user.ID,
	})
	if err != nil {
		return "", err
	}

	return "", nil
}

func (u *UserService) UpdatePassword(ctx *gin.Context, param model.UpdatePassword) (string, error) {
	user, err := helper.GetLoginUser(ctx)
	if err != nil {
		return "", err
	}

	err = helper.ComparePassword(user.Password, param.OldPassword)
	if err != nil {
		return "", err
	}

	hashPassword, err := helper.HashPassword(param.NewPassword)
	if err != nil {
		return "", err
	}

	err = u.ur.UpdatePassword(model.UpdatePassword{
		ID:              user.ID,
		OldPassword:     param.OldPassword,
		NewPassword:     hashPassword,
		ConfirmPassword: param.ConfirmPassword,
	})
	if err != nil {
		return "", err
	}

	return "", nil
}