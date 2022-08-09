package routes

import (
	"go-blog/api"
	"go-blog/middleware"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	e := gin.Default()
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := e.Group("api/v1")
	// TODO: config swagger
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "pong",
			})
		})
		v1.GET("articles", api.Articles)
		v1.GET("articles/:id", api.ShowArticle)
		authed := v1.Group("/") //需要登陆保护
		authed.Use(middleware.JWT())
		{
			v1.POST("articles", api.CreateArticle)
			v1.PUT("articles/:id", api.UpdateArticle)
			v1.DELETE("articles/:id", api.DeleteArticle)
		}

		v1.POST("admin/login", api.AdminLogin)

	}

	// GET /articles  get articles list. query(limit, page, category, tag)
	// POST /articles create article
	// GET /articles/id  get articles detail
	// PUT /articles/id update article
	// DELETE /articles/id delete article

	// GET /tags get tags list
	// GET /categories get categories list

	return e
}
