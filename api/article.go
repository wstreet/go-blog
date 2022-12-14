package api

import (
	"go-blog/service"
	"go-blog/utils"

	"github.com/gin-gonic/gin"
)

// Articles @Tags Articles
// @Summary 获取文章列表
// @Produce json
// @Accept json
// @Param data body service.UserService true "用户名, 密码"
// @Success 200 {object} serializer.ResponseUser "{"status":200,"data":{},"msg":"ok"}"
// @Failure 500  {object} serializer.ResponseUser "{"status":500,"data":{},"Msg":{},"Error":"error"}"
// @Router /user/register [post]
func Articles(c *gin.Context) {
	listService := service.ListService{
		Tag: c.Query("tag"),
	}
	res := listService.List()
	c.JSON(200, res)
}

/**
 * 文章详情
 */
func ShowArticle(c *gin.Context) {
	showService := service.ShowService{}
	res := showService.Show(c.Param("id"))
	c.JSON(200, res)
}

/**
 * 创建文章
 */
func CreateArticle(c *gin.Context) {
	createService := service.CreateArticleService{}
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Info(err)
	}
}
func UpdateArticle(c *gin.Context) {

}
func DeleteArticle(c *gin.Context) {

}
