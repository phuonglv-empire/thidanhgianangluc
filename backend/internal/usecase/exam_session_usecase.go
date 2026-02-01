package usecase

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/phuonglv-empire/thidanhgianangluc/backend/internal/domain"
	"github.com/phuonglv-empire/thidanhgianangluc/backend/internal/repository"
)

type ExamSessionUseCase interface {
	StartExamSession(userID string, examID uint) (*domain.ExamSession, error)
	GetExamSession(id uint) (*domain.ExamSession, error)
	SubmitAnswer(sessionID uint, questionID uint, userAnswer interface{}) (*domain.Answer, error)
	SubmitExam(sessionID uint) (*ExamSessionResult, error)
	GetExamResults(sessionID uint) (*ExamSessionResult, error)
}

type ExamSessionResult struct {
	Session         *domain.ExamSession `json:"session"`
	Score           float64             `json:"score"`
	TotalPoints     int                 `json:"total_points"`
	CorrectAnswers  int                 `json:"correct_answers"`
	TotalQuestions  int                 `json:"total_questions"`
}

type examSessionUseCase struct {
	sessionRepo repository.ExamSessionRepository
	examRepo    repository.ExamRepository
	answerRepo  repository.AnswerRepository
}

// NewExamSessionUseCase creates a new exam session use case
func NewExamSessionUseCase(
	sessionRepo repository.ExamSessionRepository,
	examRepo repository.ExamRepository,
	answerRepo repository.AnswerRepository,
) ExamSessionUseCase {
	return &examSessionUseCase{
		sessionRepo: sessionRepo,
		examRepo:    examRepo,
		answerRepo:  answerRepo,
	}
}

func (u *examSessionUseCase) StartExamSession(userID string, examID uint) (*domain.ExamSession, error) {
	// Get exam to validate it exists
	exam, err := u.examRepo.GetByID(examID)
	if err != nil {
		return nil, fmt.Errorf("exam not found: %w", err)
	}

	if !exam.IsActive {
		return nil, fmt.Errorf("exam is not active")
	}

	// Create new session
	session := &domain.ExamSession{
		UserID:    userID,
		ExamID:    examID,
		StartTime: time.Now(),
		Status:    domain.SessionInProgress,
	}

	if err := u.sessionRepo.Create(session); err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	// Load the exam relationship
	session.Exam = exam

	return session, nil
}

func (u *examSessionUseCase) GetExamSession(id uint) (*domain.ExamSession, error) {
	session, err := u.sessionRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("session not found: %w", err)
	}
	return session, nil
}

func (u *examSessionUseCase) SubmitAnswer(sessionID uint, questionID uint, userAnswer interface{}) (*domain.Answer, error) {
	// Get session
	session, err := u.sessionRepo.GetByID(sessionID)
	if err != nil {
		return nil, fmt.Errorf("session not found: %w", err)
	}

	if session.Status != domain.SessionInProgress {
		return nil, fmt.Errorf("session is not in progress")
	}

	// Convert user answer to JSON
	answerJSON, err := json.Marshal(userAnswer)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal answer: %w", err)
	}

	// Check if answer already exists
	answers, _ := u.answerRepo.GetBySessionID(sessionID)
	var existingAnswer *domain.Answer
	for i := range answers {
		if answers[i].QuestionID == questionID {
			existingAnswer = &answers[i]
			break
		}
	}

	if existingAnswer != nil {
		// Update existing answer
		existingAnswer.UserAnswer = answerJSON
		if err := u.answerRepo.Update(existingAnswer); err != nil {
			return nil, fmt.Errorf("failed to update answer: %w", err)
		}
		return existingAnswer, nil
	}

	// Create new answer
	answer := &domain.Answer{
		SessionID:  sessionID,
		QuestionID: questionID,
		UserAnswer: answerJSON,
	}

	if err := u.answerRepo.Create(answer); err != nil {
		return nil, fmt.Errorf("failed to create answer: %w", err)
	}

	return answer, nil
}

func (u *examSessionUseCase) SubmitExam(sessionID uint) (*ExamSessionResult, error) {
	// Get session with answers
	session, err := u.sessionRepo.GetWithAnswers(sessionID)
	if err != nil {
		return nil, fmt.Errorf("session not found: %w", err)
	}

	if session.Status != domain.SessionInProgress {
		return nil, fmt.Errorf("session is not in progress")
	}

	// Get exam with questions and correct answers
	exam, err := u.examRepo.GetWithQuestions(session.ExamID)
	if err != nil {
		return nil, fmt.Errorf("exam not found: %w", err)
	}

	// Grade the answers
	correctCount := 0
	totalPoints := 0
	earnedPoints := 0

	for _, question := range exam.Questions {
		totalPoints += question.Points

		// Find user's answer for this question
		var userAnswer *domain.Answer
		for i := range session.Answers {
			if session.Answers[i].QuestionID == question.ID {
				userAnswer = &session.Answers[i]
				break
			}
		}

		if userAnswer != nil {
			// Compare answers (simplified comparison)
			isCorrect := compareAnswers(userAnswer.UserAnswer, question.Answer)
			userAnswer.IsCorrect = &isCorrect
			
			if isCorrect {
				correctCount++
				earnedPoints += question.Points
				userAnswer.Points = &question.Points
			} else {
				zero := 0
				userAnswer.Points = &zero
			}

			// Update answer
			u.answerRepo.Update(userAnswer)
		}
	}

	// Calculate score
	var score float64
	if totalPoints > 0 {
		score = float64(earnedPoints) / float64(totalPoints) * 100
	}

	// Update session
	now := time.Now()
	session.EndTime = &now
	session.Status = domain.SessionCompleted
	session.Score = &score
	session.TotalPoints = totalPoints

	if err := u.sessionRepo.Update(session); err != nil {
		return nil, fmt.Errorf("failed to update session: %w", err)
	}

	return &ExamSessionResult{
		Session:        session,
		Score:          score,
		TotalPoints:    totalPoints,
		CorrectAnswers: correctCount,
		TotalQuestions: len(exam.Questions),
	}, nil
}

func (u *examSessionUseCase) GetExamResults(sessionID uint) (*ExamSessionResult, error) {
	session, err := u.sessionRepo.GetWithAnswers(sessionID)
	if err != nil {
		return nil, fmt.Errorf("session not found: %w", err)
	}

	if session.Status != domain.SessionCompleted {
		return nil, fmt.Errorf("exam has not been submitted yet")
	}

	// Count correct answers
	correctCount := 0
	for _, answer := range session.Answers {
		if answer.IsCorrect != nil && *answer.IsCorrect {
			correctCount++
		}
	}

	return &ExamSessionResult{
		Session:        session,
		Score:          *session.Score,
		TotalPoints:    session.TotalPoints,
		CorrectAnswers: correctCount,
		TotalQuestions: len(session.Answers),
	}, nil
}

// compareAnswers compares user answer with correct answer
func compareAnswers(userAnswer, correctAnswer json.RawMessage) bool {
	// Simple string comparison for now
	// In production, you'd want more sophisticated comparison based on question type
	return string(userAnswer) == string(correctAnswer)
}
