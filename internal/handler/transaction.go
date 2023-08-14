package handler

import (
	"errors"
	"github.com/blazee5/finance-tracker/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *Handler) GetTransactions(c *fiber.Ctx) error {
	transaction, err := h.Service.GetTransactions(c.Params("userId"))
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(transaction)
}

func (h *Handler) CreateTransaction(c *fiber.Ctx) error {
	var input models.Transaction

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	input.UserID = c.Locals("userId").(string)

	id, err := h.Service.CreateTransaction(input)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(id)

}

func (h *Handler) GetTransaction(c *fiber.Ctx) error {
	transaction, err := h.Service.GetTransaction(c.Params("id"))

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.Status(fiber.StatusNotFound).SendString("transaction not found")
	}
	if err != nil {
		log.Infof("GetTransaction err: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if transaction.UserID != c.Locals("userId").(string) {
		return c.Status(fiber.StatusForbidden).SendString("Forbidden")
	}

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(transaction)
}

func (h *Handler) UpdateTransaction(c *fiber.Ctx) error {
	var input models.Transaction

	input.ID = c.Params("id")

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	transaction, err := h.Service.GetTransaction(c.Params("id"))

	if transaction.UserID != c.Locals("userId").(string) {
		return c.Status(fiber.StatusForbidden).SendString("Forbidden")
	}

	id, err := h.Service.UpdateTransaction(input)

	if err != nil {
		return err
	}

	if transaction.UserID != c.Locals("userId").(string) {
		return c.Status(fiber.StatusForbidden).SendString("Forbidden")
	}

	return c.Status(fiber.StatusOK).JSON(id)

}

func (h *Handler) DeleteTransaction(c *fiber.Ctx) error {
	transaction, err := h.Service.GetTransaction(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("transaction not found")
	}
	if c.Locals("userId") != transaction.UserID {
		return c.Status(fiber.StatusForbidden).SendString("Forbidden")
	}

	err = h.Service.DeleteTransaction(c.Params("id"))

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func (h *Handler) AnalyzeTransactions(c *fiber.Ctx) error {
	res, err := h.Service.AnalyzeTransactions(c.Locals("userId").(string))

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res[0])
}
