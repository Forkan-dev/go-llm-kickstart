package quiz

import (
	"time"
)

type Subject struct {
	ID          uint      `gorm:"primarykey"`
	Name        string    `gorm:"unique;not null"`
	Slug        string    `gorm:"type:char(36);uniqueIndex;not null"`
	Icon        string    `gorm:"type:varchar(255);nullable"`
	Description string    `gorm:"nullable"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:current_timestamp"`
	Topics      []Topic   `gorm:"foreignKey:SubjectID"`
}
