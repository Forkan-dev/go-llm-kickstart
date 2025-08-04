package quiz

import (
	"time"
)

type Question struct {
	ID          uint      `gorm:"primarykey"`
	QuizID      uint      `gorm:"nullable"`
	TopicID     uint      `gorm:"nullable"`
	SubjectID   uint      `gorm:"nullable"`
	Title       string    `gorm:"not null"`
	Slug        string    `gorm:"type:char(36);uniqueIndex;not null"`
	Description string    `gorm:"nullable"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:current_timestamp"`

	// Relationships
	Quiz    Quiz            `gorm:"foreignKey:QuizID"`
	Topic   Topic           `gorm:"foreignKey:TopicID"`
	Subject Subject `gorm:"foreignKey:SubjectID"`
}
