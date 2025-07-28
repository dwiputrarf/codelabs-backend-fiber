package routes

import (
	"github.com/dwiputrarf/codelabs-backend-fiber/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/healthcheck", handlers.HealthCheck)
}
