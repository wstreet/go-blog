package service

import (
	"encoding/json"
	"go-blog/model"
	"go-blog/serializer"
	"go-blog/utils"
)

type TagListService struct {
}

func (service *TagListService) List() serializer.Response {
	var result []string
	var total int

	model.DB.Model(model.Article{}).
		Select("tags").
		Find(&result)
	var tags []string
	for _, value := range result {
		var _tags []string
		json.Unmarshal([]byte(value), &_tags)
		tags = append(tags, _tags...)
	}
	tags = utils.RemoveRepByMap(tags)
	total = len(tags)
	return serializer.BuildListResponse(serializer.BuildTags(tags), uint(total))
}
