package admin

import "learning-companion/internal/model"

type Repository interface {
	FindAll() ([]*model.Admin, error)
	FindByID(id uint) (*model.Admin, error)
	Create(admin *model.Admin) error
	Update(admin *model.Admin) error
	Delete(id uint) error
}
