package http

import (
	"github.com/phuonglv-empire/thidanhgianangluc/backend/internal/repository"
	"gorm.io/gorm"
)

// NewExamRepository creates exam repository (shortcut for routes)
func NewExamRepository(db *gorm.DB) repository.ExamRepository {
	return repository.NewExamRepository(db)
}

// NewExamSessionRepository creates exam session repository (shortcut for routes)
func NewExamSessionRepository(db *gorm.DB) repository.ExamSessionRepository {
	return repository.NewExamSessionRepository(db)
}

// NewAnswerRepository creates answer repository (shortcut for routes)
func NewAnswerRepository(db *gorm.DB) repository.AnswerRepository {
	return repository.NewAnswerRepository(db)
}
