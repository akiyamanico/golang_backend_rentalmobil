package controllers

import (
	"backend_test/config"
	"backend_test/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	configInstance := config.Build()
	var users []models.User
	err := models.GetAllUsers(configInstance.DB, &users)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, users)
}
func GetUserByID(c *gin.Context) {
	configInstance := config.Build()
	userID := c.Param("id")
	var user models.User
	err := models.GetUserByID(configInstance.DB, &user, userID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	configInstance := config.Build()
	var users models.User
	log.Println(users.ID)
	log.Println(users.Username)
	log.Println(users.Password)
	log.Println(users.Level)
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Kesalahan! ": err.Error()})
		return

	}
	err := models.CreateUser(configInstance.DB, &users)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, users)
}
