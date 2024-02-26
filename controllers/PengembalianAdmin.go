package controllers

import (
	"backend_test/config"
	"backend_test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PengembalianConfirmAdmin(c *gin.Context) {
	configInstance := config.Build()
	id := c.Param("id")
	var confirm models.StatusRental
	var status models.StatusRental
	c.BindJSON(&confirm)
	err := models.PengembalianAdmin(configInstance.DB, &confirm, &status, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pengembalian Selesai!"})
}
