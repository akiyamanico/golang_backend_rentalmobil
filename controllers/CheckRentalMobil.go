package controllers

import (
	"backend_test/config"
	"backend_test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMobil(c *gin.Context) {
	configInstance := config.Build()
	var mobil []models.StatusListRental
	err := models.GetMobil(configInstance.DB, &mobil)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, mobil)
}
