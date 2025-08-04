package auth

import "time"

type RefreshToken struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"user_id" gorm:"not null;type:varchar(255)"`
	Token     string    `json:"token" gorm:"not null;type:longtext"`
	Revoked   bool      `json:"revoked" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
}
