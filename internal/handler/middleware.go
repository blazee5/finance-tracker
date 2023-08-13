package handler

import (
	"github.com/blazee5/finance-tracker/internal/service"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func (h *Handler) userIdentity(c *fiber.Ctx) error {
	header := c.Get("Authorization")

	if header == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("empty authorization header")
	}

	headerParts := strings.Fields(header)
	if len(headerParts) != 2 {
		return c.Status(fiber.StatusUnauthorized).SendString("invalid authorization header")
	}

	userId, err := service.ParseToken(headerParts[1])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	c.Locals("userId", userId)

	return c.Next()
}
