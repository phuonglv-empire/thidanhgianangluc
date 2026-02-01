package repository

import (
	"github.com/phuonglv-empire/thidanhgianangluc/backend/internal/domain"
	"gorm.io/gorm"
)

type examRepository struct {
	db *gorm.DB
}

// NewExamRepository creates a new exam repository
func NewExamRepository(db *gorm.DB) ExamRepository {
	return &examRepository{db: db}
}

func (r *examRepository) GetAll() ([]domain.Exam, error) {
	var exams []domain.Exam
	err := r.db.Where("is_active = ?", true).Order("created_at DESC").Find(&exams).Error
	return exams, err
}

func (r *examRepository) GetByID(id uint) (*domain.Exam, error) {
	var exam domain.Exam
	err := r.db.First(&exam, id).Error
	if err != nil {
		return nil, err
	}
	return &exam, nil
}

func (r *examRepository) GetWithQuestions(id uint) (*domain.Exam, error) {
	var exam domain.Exam
	err := r.db.Preload("Questions", func(db *gorm.DB) *gorm.DB {
		return db.Order("\"order\" ASC")
	}).First(&exam, id).Error
	if err != nil {
		return nil, err
	}
	return &exam, nil
}

func (r *examRepository) Create(exam *domain.Exam) error {
	return r.db.Create(exam).Error
}

func (r *examRepository) Update(exam *domain.Exam) error {
	return r.db.Save(exam).Error
}

func (r *examRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Exam{}, id).Error
}
