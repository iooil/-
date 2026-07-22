package model

import "time"

type User struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	StudentNo string    `gorm:"column:student_no;size:30;unique;not null" json:"student_no"`
	Name      string    `gorm:"size:50;not null" json:"name"`
	Major     string    `gorm:"size:100;not null" json:"major"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	Role      string    `gorm:"size:20;not null;default:student" json:"role"`
	Status    int8      `gorm:"not null;default:1" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
