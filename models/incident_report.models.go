package models

import "time"

type IncidentReport struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserID uint `gorm:"column:user_id;not null" json:"user_id"`

	LoaiSuCo string `gorm:"column:loai_su_co;not null" json:"loai_su_co"`

	MucDo string `gorm:"column:muc_do;not null" json:"muc_do"`

	TieuDe string `gorm:"column:tieu_de;not null" json:"tieu_de"`

	Latitude  float64 `gorm:"column:latitude" json:"latitude"`
	Longitude float64 `gorm:"column:longitude" json:"longitude"`

	DiaChi string `gorm:"column:dia_chi;not null" json:"dia_chi"`

	MoTaChiTiet string `gorm:"column:mo_ta_chi_tiet" json:"mo_ta_chi_tiet"`

	ImageURL string `gorm:"column:image_url" json:"image_url"`

	TrangThai string `gorm:"column:trang_thai;default:Cho xu ly" json:"trang_thai"`

	NgayTao time.Time `gorm:"column:ngay_tao" json:"ngay_tao"`
}

func (IncidentReport) TableName() string {
	return "incident_reports"
}