package params

import (
	"time"
)

type Response struct {
	Status         int         `json:"status,omitempty"`
	Message        string      `json:"message,omitempty"`
	Error          string      `json:"error,omitempty"`
	AdditionalInfo interface{} `json:"additional_info,omitempty"`
	Payload        interface{} `json:"payload,omitempty"`
}

type UserResponse struct {
	ID        int        `json:"id,omitempty"`
	Email     string     `json:"email,omitempty"`
	Username  string     `json:"username,omitempty"`
	Age       int        `json:"age,omitempty"`
	Token     string     `json:"token,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Message   string     `json:"message,omitempty"`
}

type PhotoResponse struct {
	ID        int        `json:"id,omitempty"`
	Title     string     `json:"title,omitempty"`
	Caption   string     `json:"caption,omitempty"`
	PhotoUrl  string     `json:"photo_url,omitempty"`
	UserID    int        `json:"user_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	User      *UserPhoto `json:"User,omitempty"`
}

type UserPhoto struct {
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
}

type CommentResponse struct {
	ID        int            `json:"id,omitempty"`
	Message   string         `json:"message,omitempty"`
	UserID    int            `json:"user_id,omitempty"`
	PhotoID   int            `json:"photo_id,omitempty"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	User      *UserComment   `json:"User,omitempty"`
	Photo     *PhotoResponse `json:"Photo,omitempty"`
}

type UserComment struct {
	ID       int    `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
}

type SocialMediaResponse struct {
	ID             int           `json:"id,omitempty"`
	Name           string        `json:"name,omitempty"`
	SocialMediaUrl string        `json:"social_media_url,omitempty"`
	UserID         int           `json:"user_id"`
	UpdatedAt      *time.Time    `json:"updated_at,omitempty"`
	CreatedAt      *time.Time    `json:"created_at,omitempty"`
	User           *UserResponse `json:"User,omitempty"`
}

type Socmed struct {
	SocialMedias []SocialMediaResponse `json:"social_medias"`
}
