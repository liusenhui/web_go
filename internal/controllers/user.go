package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_go/internal/models"
	"web_go/internal/services"
)

type UserController struct {
	userService services.UserServiceIF
}

var user = &models.User{}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) CreateCtrl(c *gin.Context) {
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	err := uc.userService.CreateService(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 这里应该加入更多的验证逻辑，例如检查用户名是否已存在等。
	c.JSON(http.StatusOK, gin.H{"status": "User created"})
}

func (uc *UserController) SignInCtrl(c *gin.Context) {
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	if len(user.Account) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "账号不能为空"})
		return
	}
	if len(user.Password) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "密码不能为空"})
		return
	}
	user, err := uc.userService.SignInService(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (uc *UserController) GetCtrl(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not null"})
		return
	}
	user := uc.userService.GetService(id)
	c.JSON(http.StatusOK, gin.H{"data": user})
}
