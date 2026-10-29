package repositories

import (
	"project/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() ([]models.User, error)
	GetUserById(id int) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAllUser() ([]models.User, error) {
	userAll := []models.User{}

	err := r.db.Find(&userAll).Error
	if err != nil {
		return nil, err
	}

	return userAll, nil
}

func (r *userRepository) GetUserById(id int) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, id).Error

	if err != nil{
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *models.User) error {
	err := r.db.Create(user).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdateUser(user *models.User) error {
	err := r.db.Save(user).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(id int) error {
	err := r.db.Delete(&models.User{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindUserByEmail(
	email string,
) (*models.User, error) {

	var user models.User

	err := r.db.
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}