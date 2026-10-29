package dto

type CreateUserDTO struct {
	HoVaTen string `json:"hovaten" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	MatKhau string `json:"matkhau" binding:"required,min=6"`
}