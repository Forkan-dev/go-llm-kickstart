package auth

import (
	"learning-companion/internal/model/user"
)

type Service interface {
	Login(username, password string) (*user.User, string, string, error)
	Logout(accessToken string) error
}
