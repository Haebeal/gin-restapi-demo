package controllers

import (
	"gin-restapi-demo/web-service-gin/models"
	"gin-restapi-demo/web-service-gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ユーザー一覧を取得
func GetUsers(c *gin.Context) {
	// ユーザーサービスの呼び出し
	users := services.GetUsers()
	c.JSON(http.StatusOK, users)
}

// ユーザーを作成
func CreateUser(c *gin.Context) {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	err = services.CreateUser(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Sarver Error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// ユーザーを取得
func GetUser(c *gin.Context) {

}

// ユーザーを更新
func UpdateUser(c *gin.Context) {

}

// ユーザーを削除
func DeleteUser(c *gin.Context) {

}
