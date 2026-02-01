package repository

import (
	"github.com/phuonglv-empire/thidanhgianangluc/backend/internal/domain"
	"gorm.io/gorm"
)

type answerRepository struct {
	db *gorm.DB
}

// NewAnswerRepository creates a new answer repository
func NewAnswerRepository(db *gorm.DB) AnswerRepository {
	return &answerRepository{db: db}
}

func (r *answerRepository) Create(answer *domain.Answer) error {
	return r.db.Create(answer).Error
}

func (r *answerRepository) GetByID(id uint) (*domain.Answer, error) {
	var answer domain.Answer
	err := r.db.Preload("Question").First(&answer, id).Error
	if err != nil {
		return nil, err
	}
	return &answer, nil
}

func (r *answerRepository) GetBySessionID(sessionID uint) ([]domain.Answer, error) {
	var answers []domain.Answer
	err := r.db.Where("session_id = ?", sessionID).Preload("Question").Find(&answers).Error
	return answers, err
}

func (r *answerRepository) Update(answer *domain.Answer) error {
	return r.db.Save(answer).Error
}

func (r *answerRepository) DeleteBySessionID(sessionID uint) error {
	return r.db.Where("session_id = ?", sessionID).Delete(&domain.Answer{}).Error
}
