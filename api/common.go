package api

import (
	"encoding/json"
	"fmt"
	"go-blog/serializer"
	"gopkg.in/go-playground/validator.v8"
)

//返回错误信息 ErrorResponse
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := fmt.Sprintf("Field.%s", e.Field)
			// field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := fmt.Sprintf("Tag.Valid.%s", e.Tag)
			// tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serializer.Response{
				Code:  40001,
				Msg:   fmt.Sprintf("%s%s", field, tag),
				Error: fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Code:  40001,
			Msg:   "JSON类型不匹配",
			Error: fmt.Sprint(err),
		}
	}
	return serializer.Response{
		Code:  40001,
		Msg:   "参数错误",
		Error: fmt.Sprint(err),
	}
}
