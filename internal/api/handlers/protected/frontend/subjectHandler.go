package frontend

import (
	"learning-companion/internal/response"
	"learning-companion/internal/service/subject"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubjectHandler struct {
	subjectService subject.Service
}

func NewSubjectHandler(subjectService subject.Service) *SubjectHandler {
	return &SubjectHandler{subjectService: subjectService}
}

func (h *SubjectHandler) GetSubjects(c *gin.Context) {
	subjects, err := h.subjectService.GetSubjectsForFrontend()
	if err != nil {
		response.Error(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Success(c, "Subjects retrieved successfully", subjects, http.StatusOK)
}
