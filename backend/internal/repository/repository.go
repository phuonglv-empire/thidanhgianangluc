package repository

import (
	"github.com/phuonglv-empire/thidanhgianangluc/backend/internal/domain"
)

// ExamRepository defines methods for exam data access
type ExamRepository interface {
	GetAll() ([]domain.Exam, error)
	GetByID(id uint) (*domain.Exam, error)
	GetWithQuestions(id uint) (*domain.Exam, error)
	Create(exam *domain.Exam) error
	Update(exam *domain.Exam) error
	Delete(id uint) error
}

// QuestionRepository defines methods for question data access
type QuestionRepository interface {
	GetByExamID(examID uint) ([]domain.Question, error)
	GetByID(id uint) (*domain.Question, error)
	Create(question *domain.Question) error
	Update(question *domain.Question) error
	Delete(id uint) error
}

// ExamSessionRepository defines methods for exam session data access
type ExamSessionRepository interface {
	Create(session *domain.ExamSession) error
	GetByID(id uint) (*domain.ExamSession, error)
	GetWithAnswers(id uint) (*domain.ExamSession, error)
	Update(session *domain.ExamSession) error
	GetByUserID(userID string) ([]domain.ExamSession, error)
}

// AnswerRepository defines methods for answer data access
type AnswerRepository interface {
	Create(answer *domain.Answer) error
	GetByID(id uint) (*domain.Answer, error)
	GetBySessionID(sessionID uint) ([]domain.Answer, error)
	Update(answer *domain.Answer) error
	DeleteBySessionID(sessionID uint) error
}
