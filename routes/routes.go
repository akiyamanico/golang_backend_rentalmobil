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
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
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
		grp1.PUT("/pengembalian/:id", controllers.PengembalianConfirm)
	}
	grp2 := r.Group("/admin")
	{
		grp2.GET("/listavailable", controllers.GetMobil)
		grp2.GET("/listrental/:id", controllers.GetStatusRental)
		grp2.PUT("/confirm/:id", controllers.UpdateConfirm)
		grp2.PUT("/confirmpengembalian/:id", controllers.PengembalianConfirmAdmin)
	}
	grp3 := r.Group("/rentalall")
	{
		grp3.GET("", controllers.GetStatusListRental)
	}
	memberRental := r.Group("/users")
	{
		memberRental.POST("/create_rentals", controllers.CreateRental)
	}
	return r
}
