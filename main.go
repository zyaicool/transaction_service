package main

import (
	"log"
	"os"
	"transaction-service/config"
	"transaction-service/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.LoadConfig()
	// log.Printf("Loaded config: %+v", cfg)

	app := fiber.New()
	routes.SetupRoutes(app)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
