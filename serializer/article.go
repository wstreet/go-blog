package serializer

import (
	"go-blog/model"
)

// swagger:response Resp
type Article struct {
	ID      uint   `json:"id" example:"1"`       // 任务ID
	Title   string `json:"title" example:"吃饭"`   // 题目
	Content string `json:"content" example:"睡觉"` // 内容
	View    uint64 `json:"view" example:"32"`    // 浏览量
	Status  int    `json:"status" example:"0"`
}

func BuildArticle(item model.Article) Article {
	return Article{
		ID:      item.ID,
		Title:   item.Title,
		Content: item.Content,
		Status:  item.Status,
	}
}

func BuildArticles(items []model.Article) (tasks []Article) {
	for _, item := range items {
		task := BuildArticle(item)
		tasks = append(tasks, task)
	}
	return tasks
}
