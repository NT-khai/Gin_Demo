package models

import "time"

type User struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	HoVaTen       string    `gorm:"column:hovaten;not null" json:"hovaten"`
	Email         string    `gorm:"unique;not null" json:"email"`
	MatKhau       string    `gorm:"column:matkhau" json:"-"`
	GoogleID      string    `gorm:"column:google_id" json:"google_id"`
	Avatar        string    `gorm:"column:avatar" json:"avatar"`
	Provider      string    `gorm:"default:credentials" json:"provider"`
	Role          string    `gorm:"default:user" json:"role"`
	EmailVerified bool      `gorm:"column:email_verified;default:false" json:"email_verified"`
	NgayTao       time.Time `gorm:"column:ngaytao" json:"ngaytao"`
}

func (User) TableName() string {
	return "users"
}