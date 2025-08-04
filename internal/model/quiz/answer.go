package quiz

import "time"

type Answer struct {
	ID         uint      `gorm:"primarykey"`
	QuestionID uint      `gorm:"not null"`
	Text       string    `gorm:"not null"`
	Correct    bool      `gorm:"not null;default:false"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:current_timestamp"`
	// Relationships
	Question Question `gorm:"foreignKey:QuestionID"`
}
