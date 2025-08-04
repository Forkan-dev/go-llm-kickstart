package quiz

import (
	"time"
)

type Topic struct {
	ID          uint      `gorm:"primarykey"`
	SubjectID   uint      `gorm:"not null"`
	Name        string    `gorm:"not null"`
	Slug        string    `gorm:"type:char(36);uniqueIndex;not null"`
	Description string    `gorm:"nullable"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:current_timestamp"`
	// Relationships
	Subject *Subject `gorm:"foreignKey:SubjectID"`
}
