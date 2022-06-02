package models

import (
	"fmt"
	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	BaseUser
	Username     string        `json:"username" gorm:"not null;unique" form:"username" valid:"required~Username is required"`
	Email        string        `json:"email" gorm:"not null;unique" form:"email" valid:"required~Email is required, email~Invalid format email"`
	Password     string        `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age          int           `json:"age" gorm:"not null" valid:"required~Age is required"`
	Photos       []Photo       `json:"Photos,omitempty"`
	Comments     []Comment     `json:"Comments,omitempty"`
	SocialMedias []SocialMedia `json:"SocialMedias,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	u.Password = helpers.HashPass(u.Password)
	err = errCreate
	if u.Age <= 8 {
		err = fmt.Errorf("age should greater than 8")
	}
	return
}
