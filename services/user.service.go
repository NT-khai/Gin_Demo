package services

import (
	"errors"
	"project/dto"
	"project/models"
	"project/repositories"
)

type UserService interface {
	GetAll() ([]models.User, error)
	GetUserById(id int) (*models.User, error)
	CreateUser(newUser dto.CreateUserDTO) (error)
	DeleteUser(id int) error
}

type userService struct{
	r repositories.UserRepository
}

func NewUserService(
	repo repositories.UserRepository,
) UserService {
	return &userService{
		r: repo,
	}
}
func (s *userService) GetAll() ([]models.User, error) {
	return s.r.GetAllUser()
}

func (s *userService) GetUserById(id int) (*models.User,error){
	return s.r.GetUserById(id)
}

func (s *userService) CreateUser(newUser dto.CreateUserDTO) (error){

	userExist, err := s.r.FindUserByEmail(
		newUser.Email,
	)

	if err == nil && userExist != nil {
		return errors.New("email da ton tai")
	}

	user := models.User{
		HoVaTen: newUser.HoVaTen,
		Email:   newUser.Email,
		MatKhau: newUser.MatKhau,
	}


	return s.r.CreateUser(&user)
}

func (s *userService) DeleteUser(id int) error {

	if err := s.r.DeleteUser(id); err != nil {
		return err
	}

	return nil
}