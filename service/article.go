package service

import (
	"go-blog/model"
	"go-blog/serializer"
	"go-blog/utils"
	"go-blog/utils/e"
)

//创建任务的服务
type CreateArticleService struct {
	Title      string   `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content    string   `form:"content" json:"content"`
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
	code := e.SUCCESS
	err := model.DB.Create(&article).Error
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Code:  code,
			Msg:   e.GetMsg(code),
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: code,
		Data: serializer.BuildArticle(article),
		Msg:  e.GetMsg(code),
	}
}

type ListService struct {
	Limit int `json:"limit"`
	Start int `json:"start"`
}

func (service *ListService) List() serializer.Response {
	var articles []model.Article
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}
	model.DB.Model(model.Article{}).
		Count(&total).
		Limit(service.Limit).
		Offset((service.Start - 1) * service.Limit).
		Find(&articles)
	return serializer.BuildListResponse(serializer.BuildArticles(articles), uint(total))
}
