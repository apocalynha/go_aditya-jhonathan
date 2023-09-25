package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	UserID uint   `json:"user_id" form:"user_id"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
	Judul  string `json:"judul" form:"judul"`
	Konten string `json:"konten" form:"konten"`
}
