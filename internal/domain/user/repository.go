package user

import "learning-companion/internal/model"

type Repository interface {
	FindAll() ([]*model.User, error)
	FindByID(id uint) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id uint) error
}
