package handler

import (
	"github.com/blazee5/finance-tracker/internal/domain"
	"github.com/blazee5/finance-tracker/lib/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get User
// @Tags user
// @Description get user by id
// @ID get-user
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Failure 500 {string} string "internal server error"
// @Router /api/user [post]
func (h *Handler) GetUser(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	user, err := h.service.User.GetUserById(c.Context(), userId)

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

// @Summary Update User
// @Tags user
// @Description update user
// @ID update-user
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /api/user [put]
func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	var input domain.UpdateUserRequest

	userId := c.Locals("userId").(string)

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	if err := h.validator.Struct(&input); err != nil {
		validateErr := err.(validator.ValidationErrors)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": response.ValidationError(validateErr),
		})
	}

	err := h.service.User.Update(c.Context(), userId, input)

	if err != nil {
		h.log.Infof("error while update user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "server error",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

//func (h *Handler) UploadAvatar(c *fiber.Ctx) error {
//	userId := c.Locals("userId").(string)
//
//	file, err := c.FormFile("file")
//
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"message": "file is required",
//		})
//	}
//
//	err := h.service.User
//}
