package controllers

import (
	"gin-restapi-demo/web-service-gin/models"
	"gin-restapi-demo/web-service-gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ユーザー一覧を取得
func GetUsers(c *gin.Context) {
	users := []models.User{}
	services.DbEngin.Find(&users)
	c.JSON(http.StatusOK, users)
}

// ユーザーを作成
func CreateUser(c *gin.Context) {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.DbEngin.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// ユーザーを取得
func GetUser(c *gin.Context) {
	user := models.User{}
	id := c.Param("id")

	services.DbEngin.Where("ID = ?", id).First(&user)
	c.JSON(http.StatusOK, user)
}

// ユーザーを更新
func UpdateUser(c *gin.Context) {
	user := models.User{}
	id := c.Param("id")

	// JSONのバインド
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "batRequest"})
	}

	// idが一致するレコードを更新
	services.DbEngin.Model(&user).Where("ID = ?", id).Updates(user)
	c.JSON(http.StatusOK, user)
}

// ユーザーを削除
func DeleteUser(c *gin.Context) {

}
