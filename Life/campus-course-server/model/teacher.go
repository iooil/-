package model

import "time"

type Teacher struct {
	ID           uint64    `gorm:"primaryKey" json:"id"`
	TeacherNo    string    `gorm:"column:teacher_no" json:"teacher_no"`
	Name         string    `gorm:"column:name" json:"name"`
	College      string    `gorm:"column:college" json:"college"`
	Title        string    `gorm:"column:title" json:"title"`
	Introduction string    `gorm:"column:introduction" json:"introduction"`
	AverageScore float64   `gorm:"column:average_score" json:"average_score"`
	ReviewCount  int       `gorm:"column:review_count" json:"review_count"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Teacher) TableName() string {
	return "teachers"
}
