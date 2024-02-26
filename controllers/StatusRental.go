package controllers

import (
	"backend_test/config"
	"backend_test/models"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetStatusRental(c *gin.Context) {
	configInstance := config.Build()
	var status []models.StatusRental
	err := models.GetStatusRental(configInstance.DB, &status)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, status)
}

func CreateStatusRental(c *gin.Context, status, pinjaman string) {
	configInstance := config.Build()
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 104857698)
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(form.Value)
	files := form.File["dokumenpinjaman"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}
	file := files[0]
	imageName := file.Filename
	if err := c.SaveUploadedFile(file, filepath.Join("./uploads", imageName)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var statusRental models.StatusRental
	username := c.PostForm("username")

	statusRental.Username = username
	statusRental.Pinjaman = "Menunggu Konfirmasi"
	statusRental.Status = status
	statusRental.Dokumenpinjaman = imageName
	err = models.CreateStatusRental(configInstance.DB, &statusRental)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, statusRental)
}
