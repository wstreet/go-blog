package api

import (
	"github.com/gin-gonic/gin"
	"go-blog/service"
	"go-blog/utils"
)

func UserLogin(c *gin.Context) {
	userService := service.UserService{}
	if err := c.ShouldBind(&userService); err != nil {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Info(err)
	} else {
		res := userService.Login()
		c.JSON(200, res)
	}
}

func UserRegister(c *gin.Context) {
	userService := service.UserService{}
	if err := c.ShouldBind(&userService); err != nil {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Info(err)
	} else {
		res := userService.Register()
		c.JSON(200, res)
	}
}
