package service

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/internal/repository"
	"INTERN_BCC/model"
	"INTERN_BCC/pkg/helper"
	"INTERN_BCC/pkg/supabase"
	"log"

	"github.com/google/uuid"
)

type IUserService interface {
	Register(param model.UserRegister) error
	GetUser(param model.UserParam) (entity.User, error)
	Login(param model.UserLogin) (model.UserLoginResponse, error)
	UploadPhoto(param model.UploadPhoto) (string, error)
	UpdateUser(param model.UpdateUser) (string, error)
	UpdatePassword(param model.UpdatePassword) (string, error)
	FindByID(id uuid.UUID) (entity.User, error)
}

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{
		userRepository: userRepository,
	}

}

func (us *UserService) Register(param model.UserRegister) error {
	if param.Email == "" {
		return helper.ErrorEmptyEmail()
	}
	if param.Password == "" {
		return helper.ErrorEmptyPassword()
	}
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

	_, err = us.userRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) Login(param model.UserLogin) (model.UserLoginResponse, error) {
	result := model.UserLoginResponse{}

	user, err := u.userRepository.GetUser(model.UserParam{
		Email: param.Email,
	})
	if err != nil {
		return result, err
	}

	err = helper.ComparePassword(user.Password, param.Password)
	if err != nil {
		return result, helper.ErrorInvalidPassword()
	}

	token, err := helper.CreateJWTToken(user.ID)
	if err != nil {
		return result, err
	}

	result.Token = token

	return result, nil
}

func (u *UserService) GetUser(param model.UserParam) (entity.User, error) {

	return u.userRepository.GetUser(param)
}

func (u *UserService) UploadPhoto(param model.UploadPhoto) (string, error) {
	paramUser := model.UserParam{}
	paramUser.ID = param.ID
	user, err := u.userRepository.GetUser(paramUser)
	if err != nil {
		return "", nil
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

	err = u.userRepository.UpdateUser(entity.User{
		PhotoLink: link,
	}, model.UserParam{
		ID: user.ID,
	})
	if err != nil {
		return "", err
	}

	return link, nil
}

func (u *UserService) UpdateUser(params model.UpdateUser) (string, error) {
	param := model.UserParam{}
	param.ID = params.ID

	user, err := u.userRepository.GetUser(param)
	if err != nil {
		return "", err
	}
	if params.HP != "" {
		user.HP = params.HP
	}
	if params.Name != "" {
		user.Name = params.Name
	}
	log.Println(params.HP)
	err = u.userRepository.UpdateUser(user, param)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (u *UserService) UpdatePassword(param model.UpdatePassword) (string, error) {
	paramUser := model.UserParam{
		ID: param.ID,
	}
	user, err := u.userRepository.GetUser(paramUser)
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

	err = u.userRepository.UpdatePassword(model.UpdatePassword{
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

func (u *UserService) FindByID(id uuid.UUID) (entity.User, error) {
	user, err := u.userRepository.FindByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}
