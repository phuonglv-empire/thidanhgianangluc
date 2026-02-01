package domain

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// QuestionType represents the type of question
type QuestionType string

const (
	MultipleChoice QuestionType = "multiple_choice"
	Essay          QuestionType = "essay"
	TrueFalse      QuestionType = "true_false"
)

// Question represents a question in an exam
type Question struct {
	ID          uint             `gorm:"primarykey" json:"id"`
	ExamID      uint             `gorm:"not null;index" json:"exam_id"`
	ContentHTML string           `gorm:"type:text;not null" json:"content_html"`
	Type        QuestionType     `gorm:"type:varchar(50);not null" json:"type"`
	Options     datatypes.JSON   `gorm:"type:jsonb" json:"options"`         // For multiple choice options
	Answer      datatypes.JSON   `gorm:"type:jsonb" json:"answer"`          // Correct answer(s)
	Explanation string           `gorm:"type:text" json:"explanation"`      // Answer explanation
	Points      int              `gorm:"default:1" json:"points"`           // Points for this question
	Order       int              `gorm:"default:0" json:"order"`            // Question order in exam
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	DeletedAt   gorm.DeletedAt   `gorm:"index" json:"-"`
}

// TableName specifies the table name for Question
func (Question) TableName() string {
	return "questions"
}
