package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title      string `gorm:"index;not null"`
	Content    string `gorm:"type:longtext"`
	Status     int    `gorm:"default:0"`
	Tags       datatypes.JSON
	Categories datatypes.JSON
	User       User `gorm:"ForeignKey:UserId"`
	UserId     int
}
