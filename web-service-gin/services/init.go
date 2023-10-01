package services

import (
	"gin-restapi-demo/web-service-gin/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbEngin *gorm.DB

func Init() {
	dsn := "host=localhost port=5432 user=postgres password=root dbname=grom_db sslmode=disable"
	var err error
	DbEngin, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Userテーブルを自動作成
	err = DbEngin.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
}
