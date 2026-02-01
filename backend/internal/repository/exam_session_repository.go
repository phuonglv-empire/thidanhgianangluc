package repository

import (
	"github.com/phuonglv-empire/thidanhgianangluc/backend/internal/domain"
	"gorm.io/gorm"
)

type examSessionRepository struct {
	db *gorm.DB
}

// NewExamSessionRepository creates a new exam session repository
func NewExamSessionRepository(db *gorm.DB) ExamSessionRepository {
	return &examSessionRepository{db: db}
}

func (r *examSessionRepository) Create(session *domain.ExamSession) error {
	return r.db.Create(session).Error
}

func (r *examSessionRepository) GetByID(id uint) (*domain.ExamSession, error) {
	var session domain.ExamSession
	err := r.db.Preload("Exam").First(&session, id).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *examSessionRepository) GetWithAnswers(id uint) (*domain.ExamSession, error) {
	var session domain.ExamSession
	err := r.db.Preload("Exam").Preload("Answers").Preload("Answers.Question").First(&session, id).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *examSessionRepository) Update(session *domain.ExamSession) error {
	return r.db.Save(session).Error
}

func (r *examSessionRepository) GetByUserID(userID string) ([]domain.ExamSession, error) {
	var sessions []domain.ExamSession
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&sessions).Error
	return sessions, err
}
