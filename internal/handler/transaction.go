package handler

import (
	"errors"
	"github.com/blazee5/finance-tracker/internal/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
)

// @Summary Get transactions
// @Description Get transactions
// @Tags transactions
// @Accept json
// @Produce json
// @Authorization BearerAuth "Authorization"
// @Success 200 {object} []models.Transaction
// @Router /api/transactions [get]
func (h *Handler) GetTransactions(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	transaction, err := h.service.Transaction.GetTransactions(c.Context(), userId)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(transaction)
}

// @Summary Create transaction
// @Description Create transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Authorization BearerAuth "Authorization"
// @Param transaction body models.Transaction true "Transaction"
// @Success 201 {object} string
// @Router /api/transactions [post]
func (h *Handler) CreateTransaction(c *fiber.Ctx) error {
	var input domain.Transaction

	userId := c.Locals("userId").(string)

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	id, err := h.service.Transaction.CreateTransaction(c.Context(), userId, input)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(id)

}

// @Summary Get transaction
// @Description Get transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Authorization BearerAuth "Authorization"
// @Param id path string true "Transaction ID"
// @Success 200 {object} models.Transaction
// @Router /api/transaction/{id} [get]
func (h *Handler) GetTransaction(c *fiber.Ctx) error {
	id := c.Params("id")

	transaction, err := h.service.Transaction.GetTransaction(c.Context(), id)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.Status(fiber.StatusNotFound).SendString("transaction not found")
	}

	if err != nil {
		log.Infof("GetTransaction err: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if transaction.User.ID.Hex() != c.Locals("userId").(string) {
		return c.Status(fiber.StatusForbidden).SendString("Forbidden")
	}

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(transaction)
}

// @Summary Update transaction
// @Description Update transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Authorization BearerAuth "Authorization"
// @Param id path string true "Transaction ID"
// @Param transaction body models.Transaction true "Transaction"
// @Success 200 {object} string
// @Router /api/transactions/{id} [put]
func (h *Handler) UpdateTransaction(c *fiber.Ctx) error {
	var input domain.Transaction

	id := c.Params("id")
	userId := c.Locals("userId").(string)

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	transaction, err := h.service.Transaction.GetTransaction(c.Context(), id)

	if userId != transaction.User.ID.Hex() {
		return c.Status(fiber.StatusForbidden).SendString("Forbidden")
	}

	_, err = h.service.Transaction.UpdateTransaction(c.Context(), id, input)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "server error",
		})
	}

	return c.SendStatus(fiber.StatusOK)

}

// @Summary Delete transaction
// @Description Delete transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Authorization BearerAuth "Authorization"
// @Param id path string true "Transaction ID"
// @Success 200 {object} string
// @Router /api/transactions/{id} [delete]
func (h *Handler) DeleteTransaction(c *fiber.Ctx) error {
	id := c.Params("id")
	userId := c.Locals("userId").(string)

	transaction, err := h.service.Transaction.GetTransaction(c.Context(), id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("transaction not found")
	}

	if userId != transaction.User.ID.Hex() {
		return c.Status(fiber.StatusForbidden).SendString("Forbidden")
	}

	err = h.service.Transaction.DeleteTransaction(c.Context(), id)

	if err != nil {
		h.log.Infof("error while delete transaction: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

// @Summary Analyze transactions
// @Description Analyze transactions
// @Tags transactions
// @Accept json
// @Produce json
// @Authorization BearerAuth "Authorization"
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /api/transactions/analyze [get]
func (h *Handler) AnalyzeTransactions(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	res, err := h.service.Transaction.AnalyzeTransactions(c.Context(), userId)

	if err != nil {
		h.log.Infof("error while get analyze: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
