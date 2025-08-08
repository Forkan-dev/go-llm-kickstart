package frontend

import (
	"learning-companion/internal/response"
	"learning-companion/internal/service/subject"
	"log"
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

	var sTypePtr *string

	if sType, hasQ := c.GetQuery("type"); hasQ {
		sTypePtr = &sType // set pointer to value
	} else {
		sTypePtr = nil // explicitly set nil
	}
	log.Println(sTypePtr)
	subjects, err := h.subjectService.GetSubjectsForFrontend(sTypePtr)
	if err != nil {
		response.Error(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Success(c, "Subjects retrieved successfully", subjects, http.StatusOK)
}
