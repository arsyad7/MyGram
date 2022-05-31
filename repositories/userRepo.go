package repositories

import (
	"fmt"
	"mygram/helpers"
	"mygram/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	Register(user *models.User) error
	Login(user *models.User) (*models.User, int32, error)
	FindUserByEmail(email string) (*models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db}
}

func (u *userRepo) Register(user *models.User) error {
	return u.db.Create(user).Error
}

func (u *userRepo) FindUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		msg := fmt.Sprintf("User with email %s doesnt exist", email)
		return nil, fmt.Errorf(msg)
	}

	return &user, nil
}

func (u *userRepo) Login(payload *models.User) (*models.User, int32, error) {
	var user models.User

	err := u.db.Where("email = ?", payload.Email).First(&user).Error
	if err != nil {
		msg := fmt.Sprintf("User with email %s doesnt exist", payload.Email)
		return nil, 404, fmt.Errorf(msg)
	}

	checkPass := helpers.ComparePass(user.Password, payload.Password)
	if !checkPass {
		msg := fmt.Sprintln("Wrong email / password")
		return nil, 401, fmt.Errorf(msg)
	}

	return &user, 200, nil
}
