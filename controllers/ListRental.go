package controllers

import (
	"backend_test/config"
	"backend_test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatusListRental(c *gin.Context) {
	configInstance := config.Build()
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "5")
	statusListRental, err := models.GetStatusListRental(configInstance.DB, page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, statusListRental)
}
