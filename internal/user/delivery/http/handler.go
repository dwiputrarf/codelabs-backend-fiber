package http

import (
	"codelabs-backend-fiber/internal/user/domain"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type UserHandler struct {
	usecase domain.UserUsecase
}

func NewUserHandler(router fiber.Router, uc domain.UserUsecase) {
	handler := &UserHandler{uc}
	router.Get("/users", handler.GetAll)
	router.Get("/users/:id", handler.GetByID)
	router.Post("/users", handler.Create)
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.usecase.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := h.usecase.GetByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	user := &domain.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password, // You should hash this!
	}

	err := h.usecase.Create(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	resp := UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
	}

	return c.Status(201).JSON(resp)
}
