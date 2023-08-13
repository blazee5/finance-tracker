package handler

import "github.com/gofiber/fiber/v2"

func InitRoutes(app *fiber.App) {
	auth := fiber.Router(app).Group("/auth")
	{
		auth.Post("/signup", SignUp)
		auth.Post("/signin", SignIn)
	}

	api := fiber.Router(app).Group("/api")
	{
		api.Get("/transactions/:userId", GetTransactions)
		api.Post("/transactions", CreateTransaction)
		api.Get("/transactions/:id", GetTransaction)
		api.Put("/transactions/:id", UpdateTransaction)
		api.Delete("/transactions/:id", DeleteTransaction)
	}
}
