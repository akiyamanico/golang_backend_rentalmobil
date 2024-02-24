package routes

import (
	"backend_test/controllers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
	})
	grp1 := r.Group("/users")
	{
		grp1.GET("", controllers.GetUsers)
		grp1.GET("/:id", controllers.GetUserByID)
		grp1.POST("", controllers.CreateUser)
	}
	grp2 := r.Group("/rental")
	{
		grp2.GET("", controllers.GetStatusRental)
	}
	grp3 := r.Group("/listrental")
	{
		grp3.GET("", controllers.GetStatusListRental)
	}
	adminRental := r.Group("/admin")
	{
		adminRental.PUT("/status_rentals/:id", controllers.UpdateStatusRental)
	}
	memberRental := r.Group("/users")
	{
		memberRental.PUT("/status_rentals/:id", controllers.UpdateStatusRental)
		memberRental.POST("/status_rentals", controllers.CreateStatusRental)
	}
	return r
}
