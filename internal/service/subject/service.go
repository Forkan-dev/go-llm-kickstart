package subject

type Service interface {
	GetSubjectsForFrontend() ([]SubjectDTO, error)
}
