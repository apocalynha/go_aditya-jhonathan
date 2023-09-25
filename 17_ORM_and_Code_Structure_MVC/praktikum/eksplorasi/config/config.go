package config

import (
	"eksplorasi/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

func InitDB() {
	username := "apocalynha"
	password := "apocalynha123"
	host := "localhost"
	port := "3306"
	name := "prioritas1"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		username,
		password,
		host,
		port,
		name,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{}, &models.Book{}, &models.Blog{})
}
