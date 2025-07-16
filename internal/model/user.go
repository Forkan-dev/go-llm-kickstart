package model

import (
	"time"
)

type User struct {
	ID         uint       `gorm:"primarykey"`
	Uuid       string     `gorm:"type:char(36);uniqueIndex;not null"`
	CreatedAt  time.Time  `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt  time.Time  `gorm:"type:timestamp;default:current_timestamp"`
	Username   string     `gorm:"unique;not null"`
	FirstName  string     `gorm:"not null"`
	LastName   string     `gorm:"not null"`
	Email      string     `gorm:"unique;not null"`
	Password   string     `gorm:"not null"`
	LastLogin  *time.Time `gorm:"type:timestamp;nullable"`
	IPAddress  string     `gorm:"nullable"`
	MACAddress string     `gorm:"nullable"`
}
