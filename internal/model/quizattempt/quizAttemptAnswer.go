package quizattempt

import (
	"learning-companion/internal/model/quiz"
	"time"
)

type QuizAttemptAnswer struct {
	ID            uint      `gorm:"primarykey"`
	QuizAttemptID uint      `gorm:"not null"`
	QuestionID    uint      `gorm:"not null"` // Foreign key to Question this is neede for getting the question text
	AnswerID      uint      `gorm:"not null"`
	CreatedAt     time.Time `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt     time.Time `gorm:"type:timestamp;default:current_timestamp"`
	// Relationships
	QuizAttempt QuizAttempt `gorm:"foreignKey:QuizAttemptID"`
	Answer      quiz.Answer      `gorm:"foreignKey:AnswerID"`
	Question    quiz.Question    `gorm:"foreignKey:QuestionID"` // This is needed for getting the question text
}
