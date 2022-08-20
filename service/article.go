package service

import (
	"go-blog/model"
	"go-blog/serializer"
	"go-blog/utils"
)

//创建任务的服务
type CreateArticleService struct {
	Title      string   `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content    string   `form:"content" json:"content" binding:"max=1000"`
	Status     int      `form:"status" json:"status"` //0 待办   1已完成
	Tags       []string `form:"tags" json:"tags"`
	Categories string   `form:"categories" json:"categories"`
	User       string   `form:"user" json:"user"`
}

func (service *CreateArticleService) Create(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	article := model.Article{
		User:    user,
		UserId:  user.ID,
		Title:   service.Title,
		Content: service.Content,
		Status:  0,
	}
	code := utils.SUCCESS
}
