package main

import (
	"gin-restapi-demo/web-service-gin/routes"
	"gin-restapi-demo/web-service-gin/services"
)

func main() {
	// データベースの初期化
	services.Init()

	// ルーターの設定
	r := routes.SetupRouter()
	r.Run(":8080")
}
