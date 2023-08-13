package main

import (
	"fmt"
	"github.com/blazee5/task-manager/internal/config"
	"github.com/blazee5/task-manager/internal/handler"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	cfg := config.LoadConfig()

	logger, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}

	defer logger.Sync()
	log := logger.Sugar()

	app := fiber.New()

	handler.InitRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))

}
