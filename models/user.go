package models

import (
	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	BaseUser
	Username string `json:"username" gorm:"not null;unique" form:"username" valid:"required~Username is required"`
	Email    string `json:"email" gorm:"not null;" form:"email" valid:"required~Email is required, email~Invalid format email"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age      int    `json:"age" gorm:"not null" valid:"required~Age is required, gt=8"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	u.Password = helpers.HashPass(u.Password)
	err = errCreate
	return
}
