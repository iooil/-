package model

import "time"

type Course struct {
	ID            uint64    `gorm:"primaryKey" json:"id"`
	CourseNo      string    `gorm:"column:course_no" json:"course_no"`
	CourseName    string    `gorm:"column:course_name" json:"course_name"`
	TeacherID     uint64    `gorm:"column:teacher_id" json:"teacher_id"`
	College       string    `gorm:"column:college" json:"college"`
	Credit        float64   `gorm:"column:credit" json:"credit"`
	CourseType    string    `gorm:"column:course_type" json:"course_type"`
	Description   string    `gorm:"column:description" json:"description"`
	Capacity      int       `gorm:"column:capacity" json:"capacity"`
	SelectedCount int       `gorm:"column:selected_count" json:"selected_count"`
	AverageScore  float64   `gorm:"column:average_score" json:"average_score"`
	ReviewCount   int       `gorm:"column:review_count" json:"review_count"`
	Status        int8      `gorm:"column:status" json:"status"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Course) TableName() string {
	return "courses"
}

type CourseSelection struct {
	ID         uint64     `gorm:"primaryKey" json:"id"`
	UserID     uint64     `gorm:"column:user_id" json:"user_id"`
	CourseID   uint64     `gorm:"column:course_id" json:"course_id"`
	Status     string     `gorm:"column:status" json:"status"`
	SelectedAt time.Time  `gorm:"column:selected_at" json:"selected_at"`
	DroppedAt  *time.Time `gorm:"column:dropped_at" json:"dropped_at"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

func (CourseSelection) TableName() string {
	return "course_selections"
}
