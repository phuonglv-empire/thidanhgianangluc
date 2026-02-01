package domain

import (
	"time"

	"gorm.io/gorm"
)

// Exam represents an exam with multiple questions
type Exam struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Title       string         `gorm:"size:255;not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Duration    int            `gorm:"not null" json:"duration"` // Duration in minutes
	ContentHTML string         `gorm:"type:text" json:"content_html"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	Questions   []Question     `gorm:"foreignKey:ExamID" json:"questions,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for Exam
func (Exam) TableName() string {
	return "exams"
}
