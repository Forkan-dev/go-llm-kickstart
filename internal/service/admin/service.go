package admin

import "learning-companion/internal/domain/admin"

type Service interface {
	FindAll() ([]*admin.Admin, error)
	FindByID(id uint) (*admin.Admin, error)
	Create(admin *admin.Admin) error
	Update(admin *admin.Admin) error
	Delete(id uint) error
}
