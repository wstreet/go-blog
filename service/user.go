package service

import (
	"go-blog/model"
	"go-blog/serializer"
	"go-blog/utils"
	"go-blog/utils/e"
)

type UserLoginService struct {
	UserName string `from :"userName" json:"user_name" binding:"required,min=3,max=15" example:"wstreet"`
	Password string `from :"password" json:"password" binding:"required,min=5,max=16" example:"Wstreet007"`
}

func (service *UserLoginService) Login() serializer.Response {
	var user model.User
	code := e.SUCCESS

	if err := model.DB.Where("name=?", service.UserName).First(&user); err != nil {
		//TODO 如果查询不到，返回相应的错误
	}

	if !user.CheckPassword(service.Password) {
		code = e.ErrorNotCompare
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}

	token, err := utils.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	return serializer.Response{
		Code: code,
		Data: serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:  e.GetMsg(code),
	}
}
