package auth

import (
	"learning-companion/internal/model"
)

type Service interface {
	Login(username, password string) (*model.User, string, string, error)
	Logout(accessToken string) error
}
