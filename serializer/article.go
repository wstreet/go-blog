package serializer

import (
	"go-blog/model"
	"time"

	"gorm.io/datatypes"
)

// swagger:response Resp
type Article struct {
	ID         uint           `json:"id" example:"1"`       // 任务ID
	Title      string         `json:"title" example:"吃饭"`   // 题目
	Content    string         `json:"content" example:"睡觉"` // 内容
	View       uint64         `json:"view" example:"32"`    // 浏览量
	Status     int            `json:"status" example:"0"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	UserName   string         `json:"user_name" example:"street"`
	Tags       datatypes.JSON `json:"tags"`
	Categories datatypes.JSON `json:"categories"`
}

func BuildArticle(item model.Article) Article {
	return Article{
		ID:         item.ID,
		Title:      item.Title,
		Content:    item.Content,
		Status:     item.Status,
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
		UserName:   item.User.Name,
		Tags:       item.Tags,
		Categories: item.Categories,
	}
}

func BuildArticles(items []model.Article) (articles []Article) {
	for _, item := range items {
		article := BuildArticle(item)
		articles = append(articles, article)
	}

	return articles
}
