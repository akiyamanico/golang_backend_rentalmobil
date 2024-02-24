package controllers

import (
	"backend_test/models"
	"database/sql"
	"fmt"
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
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := request.ID
	query := "SELECT status FROM status_list_rentals WHERE id = ?"
	row := db.QueryRow(query, id)

	var status models.CreateRental
	fmt.Println("ini id", id)
	err = row.Scan(&status.Status)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	if status.Status == "In Use" {
		c.JSON(http.StatusForbidden, gin.H{"message": "Mobil Sedang Digunakan!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Mobil Tersedia"})
	}
	c.JSON(http.StatusOK, "OK")

}
