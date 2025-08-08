package subject

type Service interface {
	GetSubjectsForFrontend(sType *string) ([]SubjectDTO, error) // sType is refer to subject type
}
