package domain

import (
	"time"

	"gorm.io/gorm"
)

// SessionStatus represents the status of an exam session
type SessionStatus string

const (
	SessionInProgress SessionStatus = "in_progress"
	SessionCompleted  SessionStatus = "completed"
	SessionAbandoned  SessionStatus = "abandoned"
)

// ExamSession represents a user's exam attempt
type ExamSession struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	UserID     string         `gorm:"size:100;not null;index" json:"user_id"` // For future user integration
	ExamID     uint           `gorm:"not null;index" json:"exam_id"`
	Exam       *Exam          `gorm:"foreignKey:ExamID" json:"exam,omitempty"`
	StartTime  time.Time      `gorm:"not null" json:"start_time"`
	EndTime    *time.Time     `json:"end_time,omitempty"`
	Status     SessionStatus  `gorm:"type:varchar(50);not null" json:"status"`
	Score      *float64       `json:"score,omitempty"`      // Final score
	TotalPoints int           `json:"total_points"`         // Total possible points
	Answers    []Answer       `gorm:"foreignKey:SessionID" json:"answers,omitempty"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for ExamSession
func (ExamSession) TableName() string {
	return "exam_sessions"
}
