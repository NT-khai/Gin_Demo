package models

import "time"

type Feedback struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	HoVaTen   string    `gorm:"column:hovaten;not null" json:"hovaten"`
	Email     string    `gorm:"column:email;not null" json:"email"`
	Content   string    `gorm:"column:content;not null" json:"content"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (Feedback) TableName() string {
	return "feedback"
}