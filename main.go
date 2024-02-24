package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"backend_test/config"
	"backend_test/routes"
)

func main() {
	configInstance := config.Build()
	defer configInstance.DB.Close()
	r := routes.SetupRouter()
	address := "localhost"
	port := "8080"
	r.Run(fmt.Sprintf("%s:%s", address, port))
	fmt.Printf("Server is running on http://%s:%s\n", address, port)
}
