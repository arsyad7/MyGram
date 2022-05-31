package params

import (
	"time"
)

type Response struct {
	Status         int         `json:"status"`
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
