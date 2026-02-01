package usecase

import (
	"fmt"

	"github.com/phuonglv-empire/thidanhgianangluc/backend/internal/domain"
	"github.com/phuonglv-empire/thidanhgianangluc/backend/internal/repository"
)

type ExamUseCase interface {
	GetAllExams() ([]domain.Exam, error)
	GetExamByID(id uint) (*domain.Exam, error)
	GetExamWithQuestions(id uint) (*domain.Exam, error)
}

type examUseCase struct {
	examRepo repository.ExamRepository
}

// NewExamUseCase creates a new exam use case
func NewExamUseCase(examRepo repository.ExamRepository) ExamUseCase {
	return &examUseCase{
		examRepo: examRepo,
	}
}

func (u *examUseCase) GetAllExams() ([]domain.Exam, error) {
	exams, err := u.examRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get exams: %w", err)
	}
	return exams, nil
}

func (u *examUseCase) GetExamByID(id uint) (*domain.Exam, error) {
	exam, err := u.examRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get exam: %w", err)
	}
	return exam, nil
}

func (u *examUseCase) GetExamWithQuestions(id uint) (*domain.Exam, error) {
	exam, err := u.examRepo.GetWithQuestions(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get exam with questions: %w", err)
	}
	
	// Don't expose correct answers to users
	if exam.Questions != nil {
		for i := range exam.Questions {
			exam.Questions[i].Answer = nil
		}
	}
	
	return exam, nil
}
