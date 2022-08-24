package api

import (
	"github.com/gin-gonic/gin"
	"go-blog/service"
	"go-blog/utils"
)

func UserLogin(c *gin.Context) {
	userLoginService := service.UserLoginService{}
	if err := c.ShouldBind(&userLoginService); err != nil {
		res := userLoginService.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Info(err)
	}
}
