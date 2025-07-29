package handlers

import "github.com/gofiber/fiber/v2"

func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Api is running from codelabs-backend-fiber!",
	})
}
