package handler

import (
	"github.com/blazee5/finance-tracker/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {
	var input models.User

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	id, err := h.Service.CreateUser(input)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(id)
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	var input models.User

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	token, err := h.Service.GenerateToken(input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON("invalid credentials")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
