package quiz

import (
	"time"
)

type Difficulty string

const (
	DifficultyEasy   Difficulty = "easy"
	DifficultyMedium Difficulty = "medium"
	DifficultyHard   Difficulty = "hard"
)

type Type string

const (
	TypeMCQ      Type = "MCQ"
	TypeMockTest Type = "MockTest"
	Timed        Type = "Timed"
)

type Quiz struct {
	ID                  uint       `gorm:"primarykey"`
	TopicID             uint       `gorm:"not null"`
	SubjectID           uint       `gorm:"not null"`
	Title               string     `gorm:"not null"`
	Slug                string     `gorm:"type:char(36);uniqueIndex;not null"`
	Description         string     `gorm:"nullable"`
	Difficulty          Difficulty `gorm:"type:enum('easy', 'medium', 'hard');not null"`
	Type                Type       `gorm:"type:enum('MCQ', 'MockTest', 'Timed');not null"`
	is_timed            bool       `gorm:"not null;default:false"`
	Duration_in_seconds int        `gorm:"not null;default:0"` // Duration in seconds
	CreatedAt           time.Time  `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt           time.Time  `gorm:"type:timestamp;default:current_timestamp"`
	// Relationships
	Subject Subject `gorm:"foreignKey:SubjectID"`
	Topic   Topic   `gorm:"foreignKey:TopicID"`
}
