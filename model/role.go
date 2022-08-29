package model

import (
	"gorm.io/gorm"
)

const (
	ADMIN = "admin"
	GUEST = "guest"
)

type Role struct {
	gorm.Model
	Type string
}

// 内置两种角色

func CreateDefaultRoles() {
	admin := Role{
		Type: ADMIN,
	}
	guest := Role{
		Type: GUEST,
	}
	roles := []Role{admin, guest}
	DB.Create(&roles)
}
