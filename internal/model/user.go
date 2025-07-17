package model

import (
	"learning-companion/pkg/jwt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
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

func (user *User) CreateToken() string {
	token, err := jwt.GenerateAccessToken(user.Uuid)
	if err != nil {
		return ""
	}

	return token
}

func (user *User) CreateRefreshToken() string {
	token, err := jwt.GenerateRefreshToken(user.Uuid)
	if err != nil {
		return ""
	}

	return token
}

func (user *User) ParseToken(tokenString string) (*jwt.TokenClaims, error) {
	claims, err := jwt.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func (user *User) CheckPassword(password string) bool {
	// Implement password checking logic here, e.g., using bcrypt

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println("Password check failed:", err)
		return false
	}

	return true
}
