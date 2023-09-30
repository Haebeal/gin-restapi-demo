package main

import (
	"gin-restapi-demo/web-service-gin/models"
	"gin-restapi-demo/web-service-gin/routes"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	dsn := "host=localhost port=5432 user=postgres password=root dbname=grom_db sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Userテーブルを自動作成
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	// DBのクローズ
	defer func() {
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	// ルーターの設定
	r := routes.SetupRouter()

	r.Run(":8080")
}
