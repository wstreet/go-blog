package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Email    string
	Password string
	RoleId   int
	Role     Role `gorm:"ForeignKey:RoleId"`
}

var PasswordDifficulty = 12

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordDifficulty)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// 内置管理员
func CreateDefaultUsers() {
	var role Role
	DB.Model(&Role{}).Where("type=?", ADMIN).First(&role)
	admin := User{
		Name: "admin",
		Role: role,
	}
	admin.SetPassword("123456")
	DB.Create(&admin) // TODO: 判断是否执行过
}
