package handler

import (
	"github.com/blazee5/finance-tracker/internal/domain"
	"github.com/blazee5/finance-tracker/lib/response"
	"github.com/go-playground/validator/v10"
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
	var input domain.SignUpRequest

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	if err := h.validator.Struct(&input); err != nil {
		validateErr := err.(validator.ValidationErrors)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": response.ValidationError(validateErr),
		})
	}

	token, err := h.service.Auth.CreateUser(c.Context(), input)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
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
	var input domain.SignInRequest

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	if err := h.validator.Struct(&input); err != nil {
		validateErr := err.(validator.ValidationErrors)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": response.ValidationError(validateErr),
		})
	}

	token, err := h.service.Auth.GenerateToken(c.Context(), input.Email, input.Password)

	if err != nil {
		h.log.Infof("error while sign in: %v", err)
		return c.Status(fiber.StatusUnauthorized).JSON("invalid credentials")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

// @Summary Get User
// @Tags auth
// @Description get user by id
// @ID get-user
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Failure 500 {string} string "internal server error"
// @Router /api/user [post]
func (h *Handler) GetUser(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	user, err := h.service.Auth.GetUserById(c.Context(), userId)

	if err != nil {
		h.log.Infof("error while get user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.User{
		ID:    user.ID.Hex(),
		Name:  user.Name,
		Email: user.Email,
	})
}
