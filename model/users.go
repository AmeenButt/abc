package models

import "time"

type ValidApi struct {
	XApiKey string `header:"x-api-key"`
}

type UserLoggedIn struct {
	Token string `header:"token"`
}

type UserProfile struct {
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	IsVerified  bool   `json:"is_verified"`
}

type UserToken struct {
	ID          int         `json:"id"`
	Username    string      `json:"username"`
	Email       string      `json:"email"`
	LastLogin   time.Time   `json:"last_loggedin_at"`
	UserProfile UserProfile `json:"user_profile"`
}
