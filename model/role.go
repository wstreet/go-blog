package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Type string
}

// 内置两种角色

func CreateDefaultRoles() {
	admin := Role{
		Type: "admin",
	}
	guest := Role{
		Type: "guest",
	}
	roles := []Role{admin, guest}
	DB.Create(&roles)
}
