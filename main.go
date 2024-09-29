package main

import (
	"eCommerceService/database"
	"eCommerceService/seeder"
	"eCommerceService/src/router"
	"fmt"
)

func main() {
	fmt.Println("Starting server on port 8080...")

	database.Migrate()
	seeder.Seeder()

	route := router.SetupRouter()
	err := route.Run(":8080")

	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
