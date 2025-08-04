package quizattempt

import (
	"learning-companion/internal/model/quiz"
	"learning-companion/internal/model/user"
	"time"
)

type QuizAttempt struct {
	ID             uint      `gorm:"primarykey"`
	QuizID         uint      `gorm:"not null"`
	UserID         uint      `gorm:"not null"`
	ScoreThreshold int       `gorm:"not null;default:0"`   // Minimum score required to pass the quiz
	FullScore      int       `gorm:"not null;default:100"` // Maximum score for the quiz
	Score          int       `gorm:"not null;default:0"`
	Attempted      bool      `gorm:"not null;default:false"` // Indicates if the quiz was attempted
	CreatedAt      time.Time `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt      time.Time `gorm:"type:timestamp;default:current_timestamp"`
	// Relationships
	Quiz quiz.Quiz `gorm:"foreignKey:QuizID"`
	User user.User `gorm:"foreignKey:UserID"`
}
