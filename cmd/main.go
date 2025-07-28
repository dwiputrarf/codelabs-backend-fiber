package main

import (
	"log"

	"github.com/dwiputrarf/codelabs-backend-fiber/config"
	"github.com/dwiputrarf/codelabs-backend-fiber/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()

	app := fiber.New()

	routes.SetupRoutes(app)

	port := config.GetEnv("PORT", "3000")
	log.Fatal(app.Listen(":" + port))
}
