package models

type UserProfile struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	UserID      uint   `gorm:"column:user_id;unique;not null" json:"user_id"`
	HoVaTen     string `gorm:"column:hovaten;not null" json:"hovaten"`
	SoDienThoai string `gorm:"column:so_dien_thoai" json:"so_dien_thoai"`
	DiaChi      string `gorm:"column:dia_chi" json:"dia_chi"`
	Avatar      string `gorm:"column:avatar" json:"avatar"`
}

func (UserProfile) TableName() string {
	return "user_profiles"
}