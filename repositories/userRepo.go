package repositories

import (
	"mygram/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	Register(user *models.User) error
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
