package model

import "time"

type RefreshToken struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"user_id" gorm:"not null"`
	Token     string    `json:"token" gorm:"not null;unique"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
}
