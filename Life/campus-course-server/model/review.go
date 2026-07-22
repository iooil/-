package model

import "time"

type MyCourseReview struct {
	ID uint64 `json:"id"`

	CourseName string `json:"course_name"`

	Score int `json:"score"`

	Content string `json:"content"`

	CreatedAt time.Time `json:"created_at"`
}

type MyTeacherReview struct {
	ID uint64 `json:"id"`

	TeacherName string `json:"teacher_name"`

	Score int `json:"score"`

	Content string `json:"content"`

	CreatedAt time.Time `json:"created_at"`
}

type CourseReview struct {
	ID uint64 `gorm:"primaryKey" json:"id"`

	UserID uint64 `json:"user_id"`

	CourseID uint64 `json:"course_id"`

	Score int `json:"score"`

	Content string `json:"content"`

	IsAnonymous int `json:"is_anonymous"`

	CreatedAt time.Time `json:"created_at"`
}

func (CourseReview) TableName() string {

	return "course_reviews"

}

type TeacherReview struct {
	ID uint64 `gorm:"primaryKey" json:"id"`

	UserID uint64 `json:"user_id"`

	TeacherID uint64 `json:"teacher_id"`

	CourseID uint64 `json:"course_id"`

	Score int `json:"score"`

	Content string `json:"content"`

	IsAnonymous int `json:"is_anonymous"`

	CreatedAt time.Time `json:"created_at"`
}

func (TeacherReview) TableName() string {

	return "teacher_reviews"

}
