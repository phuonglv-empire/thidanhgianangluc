package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phuonglv-empire/thidanhgianangluc/backend/internal/usecase"
)

type ExamHandler struct {
	examUseCase usecase.ExamUseCase
}

// NewExamHandler creates a new exam handler
func NewExamHandler(examUseCase usecase.ExamUseCase) *ExamHandler {
	return &ExamHandler{
		examUseCase: examUseCase,
	}
}

// GetExams handles GET /api/exams
func (h *ExamHandler) GetExams(c *gin.Context) {
	exams, err := h.examUseCase.GetAllExams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exams)
}

// GetExam handles GET /api/exams/:id
func (h *ExamHandler) GetExam(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid exam id"})
		return
	}

	exam, err := h.examUseCase.GetExamWithQuestions(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "exam not found"})
		return
	}

	c.JSON(http.StatusOK, exam)
}
