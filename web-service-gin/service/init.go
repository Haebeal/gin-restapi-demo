package service

import (
	"gin-restapi-demo/web-service-gin/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// テーブルの初期化
func init() {
	dsn := "user=postgres password=postgres dbname=postgres sslmode=disable"

	// PostgreSQLに接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// エラーのロギング
		log.Fatal(err.Error())
	}

	// マイグレーションの実行
	db.AutoMigrate(&models.User{})
}
