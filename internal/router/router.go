package routes

import (
	"codelabs-backend-fiber/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/healthcheck", handlers.HealthCheck)
}
