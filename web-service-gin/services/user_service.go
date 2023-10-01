package services

import (
	"gin-restapi-demo/web-service-gin/models"
)

func GetUsers() []models.User {
	var users []models.User
	DbEngin.Find(&users)
	return users
}

func CreateUser() {
}
