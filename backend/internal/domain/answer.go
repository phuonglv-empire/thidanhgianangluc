package domain

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Answer represents a user's answer to a question
type Answer struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	SessionID  uint           `gorm:"not null;index" json:"session_id"`
	QuestionID uint           `gorm:"not null;index" json:"question_id"`
	Question   *Question      `gorm:"foreignKey:QuestionID" json:"question,omitempty"`
	UserAnswer datatypes.JSON `gorm:"type:jsonb" json:"user_answer"` // User's answer
	IsCorrect  *bool          `json:"is_correct,omitempty"`          // Null for ungraded
	Points     *int           `json:"points,omitempty"`              // Points earned
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for Answer
func (Answer) TableName() string {
	return "answers"
}
