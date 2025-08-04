package http

import (
	"codelabs-backend-fiber/internal/user/domain"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, uc domain.UserUsecase) {
	api := app.Group("/api")
	NewUserHandler(api, uc)
}
