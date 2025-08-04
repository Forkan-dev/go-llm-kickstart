package domain

import "gorm.io/gorm"

// Repository defines the common interface for all repositories.
// It uses a type parameter T to represent the model type.
type Repository[T any] interface {
	Create(entity *T) error
	FindByID(id uint) (*T, error)
	Update(entity *T) error
	Delete(entity *T) error
	FindAll() ([]T, error)
	GetDB() *gorm.DB
}
