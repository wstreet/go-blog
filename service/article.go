package service

import (
	"go-blog/model"
	"go-blog/serializer"
	"go-blog/utils"
	"go-blog/utils/e"

	"gorm.io/datatypes"
)

// 创建任务的服务
type CreateArticleService struct {
	Title      string         `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content    string         `form:"content" json:"content"`
	Status     int            `form:"status" json:"status"` //0 待办   1已完成
	Tags       datatypes.JSON `form:"tags" json:"tags"`
	Categories datatypes.JSON `form:"categories" json:"categories"`
	User       string         `form:"user" json:"user"`
}

func (service *CreateArticleService) Create(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	article := model.Article{
		User:       user,
		UserId:     user.ID,
		Title:      service.Title,
		Content:    service.Content,
		Status:     0,
		Tags:       service.Tags,
		Categories: service.Categories,
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
	Limit int    `json:"limit"`
	Start int    `json:"start"`
	Tag   string `json: "tag"`
}

func (service *ListService) List() serializer.Response {
	var articles []model.Article
	var total int64
	if service.Limit == 0 {
		service.Limit = 10
	}

	model.DB.Model(model.Article{}).
		Where("tags LIKE ?", "%"+service.Tag+"%"). // TODO: 查询待修改
		Count(&total).
		Limit(service.Limit).
		Offset((service.Start - 1) * service.Limit).
		Find(&articles)
	return serializer.BuildListResponse(serializer.BuildArticles(articles), uint(total))
}

type ShowService struct {
}

func (service *ShowService) Show(id string) serializer.Response {
	var article model.Article
	var user model.User
	code := e.SUCCESS
	err := model.DB.First(&article, id).Error
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Code:  code,
			Msg:   e.GetMsg(code),
			Error: err.Error(),
		}
	}
	article.AddView()
	model.DB.First(&user, article.UserId)
	article.User = user
	return serializer.Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: serializer.BuildArticle(article),
	}
}
