package auth

import (
	"github.com/joatisio/wisp/internal/models"
)

type TokenPair struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

type LoginRequest struct {
	Email    string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	User      models.User `json:"user"`
	TokenPair TokenPair   `json:"token_pair"`
}

type LogoutRequest struct {
	AccessToken string    `json:"access_token"`
	UserID      models.ID `json:"user_id"`
}

type RegisterRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	CPassword string `json:"c_password" binding:"required"`
}

type RegisterResponse struct {
	User models.User `json:"user"`
}
