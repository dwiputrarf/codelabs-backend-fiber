package main

import (
	"log"

	"codelabs-backend-fiber/config"
	"codelabs-backend-fiber/internal/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.InitPostgreDB()

	app := fiber.New()

	routes.SetupRoutes(app)

	port := config.GetEnv("PORT", "3000")
	log.Fatal(app.Listen(":" + port))
}
