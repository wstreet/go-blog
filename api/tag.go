package api

import (
	"go-blog/service"

	"github.com/gin-gonic/gin"
)

func Tags(c *gin.Context) {
	listService := service.TagListService{}
	res := listService.List()
	c.JSON(200, res)
}
