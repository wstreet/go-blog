package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	e := gin.Default()
	v1 := e.Group("api/v1")
	// TODO: config swagger
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "pong",
		})
	})
	// GET /articles  get articles list. query(limit, page, category, tag)
	// POST /articles create article
	// GET /articles/id  get articles detail
	// PUT /articles/id update article
	// DELETE /articles/id delete article

	// GET /tags get tags list
	// GET /categories get categories list

	return e
}
