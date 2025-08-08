package subject

import (
	"learning-companion/internal/model/quiz"

	"gorm.io/gorm"
)

type serviceImpl struct {
	db *gorm.DB
}

type TopicDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

type SubjectDTO struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	Icon        string     `json:"icon"`
	Description string     `json:"description"`
	Topics      []TopicDTO `json:"topics"`
}

func NewService(db *gorm.DB) Service {
	return &serviceImpl{db: db}
}

func (s *serviceImpl) GetSubjectsForFrontend(sType *string) ([]SubjectDTO, error) {
	var subjects []quiz.Subject
	err := s.db.Limit(4).Find(&subjects).Error
	if err != nil {
		return nil, err
	}

	var subjectDTOs []SubjectDTO
	for _, subject := range subjects {
		subjectDTO := SubjectDTO{
			ID:          subject.ID,
			Name:        subject.Name,
			Slug:        subject.Slug,
			Icon:        subject.Icon,
			Description: subject.Description,
			Topics:      []TopicDTO{},
		}

		subjectDTOs = append(subjectDTOs, subjectDTO)
	}

	return subjectDTOs, nil
}
