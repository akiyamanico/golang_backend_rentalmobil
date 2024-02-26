package controllers

import (
	"backend_test/config"
	"backend_test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateConfirm(c *gin.Context) {
	configInstance := config.Build()
	id := c.Param("id")
	var confirm models.Confirm
	var status models.StatusRental

	c.BindJSON(&confirm)
	err := models.UpdateConfirm(configInstance.DB, &confirm, &status, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Terkonfirmasi"})
}
