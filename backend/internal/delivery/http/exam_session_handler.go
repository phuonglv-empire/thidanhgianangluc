package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phuonglv-empire/thidanhgianangluc/backend/internal/usecase"
)

type ExamSessionHandler struct {
	sessionUseCase usecase.ExamSessionUseCase
}

// NewExamSessionHandler creates a new exam session handler
func NewExamSessionHandler(sessionUseCase usecase.ExamSessionUseCase) *ExamSessionHandler {
	return &ExamSessionHandler{
		sessionUseCase: sessionUseCase,
	}
}

// CreateSessionRequest represents request to create a session
type CreateSessionRequest struct {
	ExamID uint   `json:"exam_id" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
}

// SubmitAnswerRequest represents request to submit an answer
type SubmitAnswerRequest struct {
	QuestionID uint        `json:"question_id" binding:"required"`
	UserAnswer interface{} `json:"user_answer" binding:"required"`
}

// CreateSession handles POST /api/exam-sessions
func (h *ExamSessionHandler) CreateSession(c *gin.Context) {
	var req CreateSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := h.sessionUseCase.StartExamSession(req.UserID, req.ExamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, session)
}

// GetSession handles GET /api/exam-sessions/:id
func (h *ExamSessionHandler) GetSession(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session id"})
		return
	}

	session, err := h.sessionUseCase.GetExamSession(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}

	c.JSON(http.StatusOK, session)
}

// SubmitAnswer handles POST /api/exam-sessions/:id/answers
func (h *ExamSessionHandler) SubmitAnswer(c *gin.Context) {
	idStr := c.Param("id")
	sessionID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session id"})
		return
	}

	var req SubmitAnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	answer, err := h.sessionUseCase.SubmitAnswer(uint(sessionID), req.QuestionID, req.UserAnswer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, answer)
}

// SubmitExam handles POST /api/exam-sessions/:id/submit
func (h *ExamSessionHandler) SubmitExam(c *gin.Context) {
	idStr := c.Param("id")
	sessionID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session id"})
		return
	}

	result, err := h.sessionUseCase.SubmitExam(uint(sessionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetResults handles GET /api/exam-sessions/:id/results
func (h *ExamSessionHandler) GetResults(c *gin.Context) {
	idStr := c.Param("id")
	sessionID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session id"})
		return
	}

	result, err := h.sessionUseCase.GetExamResults(uint(sessionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
