package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Email    string
	Password string
	RoleId   int
	Role     Role `gorm:"ForeignKey:RoleId"`
}

var PasswordDifficulty = 12
