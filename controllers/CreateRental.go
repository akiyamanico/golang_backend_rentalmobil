package controllers

import (
	"backend_test/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CreateRental(c *gin.Context) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/rentalmobil")
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	var request models.CreateRental
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idValues := form.Value["id"]
	idValue := idValues[0]
	files := form.File["dokumenpinjaman"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	file := files[0]
	query := "SELECT status, namamobil FROM status_list_rentals WHERE id = ?"
	row := db.QueryRow(query, idValue)

	imageName := file.Filename
	var status models.CreateRental
	username := c.PostForm("username")
	request.Dokumenpinjaman = imageName
	request.Username = username
	err = row.Scan(&status.Status, &status.Pinjaman)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	if status.Status == "In Use" {
		c.JSON(http.StatusForbidden, gin.H{"message": "Mobil Sedang Digunakan!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Mobil Tersedia"})
		CreateStatusRental(c, status.Status, status.Pinjaman)
	}

}
