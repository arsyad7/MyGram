package models

import "gorm.io/gorm"

type SocialMedia struct {
	Name           string `json:"name" gorm:"not null" form:"name" valid:"required~Name is required"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null" form:"social_media_url" valid:"required~Social media url is required"`
	UserID         uint   `json:"user_id"`
	User           User
	*gorm.Model
}
