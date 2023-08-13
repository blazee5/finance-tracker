package handler

import (
	"github.com/blazee5/finance-tracker/internal/service"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes(app *fiber.App) {
	auth := fiber.Router(app).Group("/auth")
	{
		auth.Post("/signup", h.SignUp)
		auth.Post("/signin", h.SignIn)
	}

	api := fiber.Router(app).Group("/api")
	{
		api.Get("/transactions/:userId", h.userIdentity, h.GetTransactions)
		api.Post("/transactions", h.userIdentity, h.CreateTransaction)
		api.Get("/transaction/:id", h.userIdentity, h.GetTransaction)
		api.Put("/transactions/:id", h.userIdentity, h.UpdateTransaction)
		api.Delete("/transactions/:id", h.userIdentity, h.DeleteTransaction)
	}
}
