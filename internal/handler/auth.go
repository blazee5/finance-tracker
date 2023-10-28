package handler

import (
	"github.com/blazee5/finance-tracker/internal/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Sign up
// @Tags auth
// @Description sign up
// @ID sign-up
// @Accept  json
// @Produce  json
// @Param user body models.User true "User"
// @Success 201 {string} string	"ok"
//
//	@Example request: {
//	  "email": "test@test",
//	  "password": "test"
//	}
//
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /auth/signup [post]
func (h *Handler) SignUp(c *fiber.Ctx) error {
	var input models.User

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	id, err := h.Service.CreateUser(input)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

// @Summary Sign in
// @Tags auth
// @Description sign in
// @ID sign-in
// @Accept  json
// @Produce  json
// @Param user body models.User true "User"
// @Success 200 {string} string	"ok"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /auth/signin [post]
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
