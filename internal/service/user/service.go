package user

import "learning-companion/internal/domain/user"

type Service interface {
	FindAll() ([]*user.User, error)
	FindByID(id uint) (*user.User, error)
	Create(user *user.User) error
	Update(user *user.User) error
	Delete(id uint) error
}
