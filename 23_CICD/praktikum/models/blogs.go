package models

import (
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	gorm.Model
	UsersID uint   `json:"users_id" form:"users_id"`
	User    User   `json:"user" gorm:"foreignkey:UsersID"` // Explicitly define the foreign key
	Judul   string `json:"judul" form:"judul"`
	Konten  string `json:"konten" form:"konten"`
}

type BlogResponse struct {
	ID        uint      `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      User      `json:"user" ` // Explicitly define the foreign key
	Judul     string    `json:"judul" `
	Konten    string    `json:"konten"`
}

func (blogDB Blog) ResponseConvert() BlogResponse {
	var Response BlogResponse
	Response.ID = blogDB.ID
	Response.Judul = blogDB.Judul
	Response.Konten = blogDB.Konten

	return Response
}
