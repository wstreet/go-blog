package serializer

import (
	"go-blog/model"
)

// swagger:response Resp
type User struct {
	ID     uint   `json:"id" form:"id" example:"1"`
	Name   string `json:"name"`
	Email  string
	RoleId int
}

func BuildUser(user model.User) User {
	return User{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		RoleId: user.RoleId,
	}
}
