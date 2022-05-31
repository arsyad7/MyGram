package repositories

import (
	"fmt"
	"mygram/helpers"
	"mygram/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo interface {
	Register(user *models.User) error
	Login(user *models.User) (*models.User, int32, error)
	FindUserByEmail(email string) (*models.User, error)
	CheckUser(id uint) error
	UpdateUser(user *models.User, id uint) (*models.User, error)
	DeleteUser(id uint) error
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

func (u *userRepo) CheckUser(id uint) error {
	return u.db.Where("id = ?", id).Error
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

func (u *userRepo) UpdateUser(p *models.User, id uint) (*models.User, error) {
	var user models.User

	err := u.db.Model(&user).Table("users").Clauses(clause.Returning{}).Where("id = ?", id).Updates(models.User{Username: p.Username, Email: p.Email}).Order("created_at ASC").Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) DeleteUser(id uint) error {
	var user models.User
	err := u.db.Where("id = ?", id).First(&user).Error
	if err == nil {
		err = u.db.Where("id = ?", id).Delete(&user).Error
	}
	return err
}
