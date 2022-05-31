package models

type Photo struct {
	Title    string `json:"title" gorm:"not null" form:"title" valid:"required~Title is required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" gorm:"not null" form:"photo_url" valid:"required~Photo URL is required"`
	UserID   uint   `json:"user_id"`
	BasePhoto
}
