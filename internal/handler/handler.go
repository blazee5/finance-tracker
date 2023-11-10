package handler

import (
	"github.com/blazee5/finance-tracker/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"go.uber.org/zap"
)

type Handler struct {
	log     *zap.SugaredLogger
	service *service.Service
}

func NewHandler(log *zap.SugaredLogger, service *service.Service) *Handler {
	return &Handler{log: log, service: service}
}

func (h *Handler) InitRoutes(app *fiber.App) {
	auth := fiber.Router(app).Group("/auth")
	{
		auth.Post("/signup", h.SignUp)
		auth.Post("/signin", h.SignIn)
	}

	api := fiber.Router(app).Group("/api")
	{
		user := api.Group("/user")
		{
			user.Get("/", h.userIdentity, h.GetUser)
		}

		transactions := api.Group("/transactions")
		{
			transactions.Get("/analyze", h.userIdentity, h.AnalyzeTransactions)
			transactions.Get("/", h.userIdentity, h.GetTransactions)
			transactions.Post("/", h.userIdentity, h.CreateTransaction)
			transactions.Get("/:id", h.userIdentity, h.GetTransaction)
			transactions.Put("/:id", h.userIdentity, h.UpdateTransaction)
			transactions.Delete("/:id", h.userIdentity, h.DeleteTransaction)
		}
	}

	app.Get("/swagger/*", swagger.HandlerDefault)
}
