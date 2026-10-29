package controllers

import (
	"net/http"
	"project/dto"
	"project/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(
	service services.UserService,
) *UserController {
	return &UserController{
		userService: service,
	}
}

func (c *UserController) GetAll(ctx *gin.Context) {
	users, err := c.userService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetUserById(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id không hợp lệ",
		})
		return
	}

	user, err := c.userService.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) CreateUser(ctx *gin.Context) {

	var body dto.CreateUserDTO

	// Bind JSON
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Du lieu khong hop le",
			"error": err.Error(),
		})
		return
	}

	// Goi Service
	err := c.userService.CreateUser(dto.CreateUserDTO(body))

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Thanh cong
	ctx.JSON(201, gin.H{
		"message": "Tao user thanh cong",
	})
}

func (c *UserController) DeleteUser(ctx *gin.Context){
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "id khong hop le",
		})
		return
	}

	err = c.userService.DeleteUser(id)

	if err != nil{
		ctx.JSON(500, gin.H{
			"error":err.Error(),
		})

		ctx.JSON(200, gin.H{
		"message": "Xoa user thanh cong",
		})
	}
}