package main

import (
	"fmt"
	_ "github.com/blazee5/finance-tracker/docs"
	"github.com/blazee5/finance-tracker/internal/config"
	"github.com/blazee5/finance-tracker/internal/handler"
	"github.com/blazee5/finance-tracker/internal/repository"
	"github.com/blazee5/finance-tracker/internal/repository/mongodb"
	"github.com/blazee5/finance-tracker/internal/repository/redis"
	"github.com/blazee5/finance-tracker/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

// @title Finance Tracker API
// @version 1.0
// @description Finance Tracker API Documentation
// @host localhost:3000
// @BasePath /
func main() {
	cfg := config.LoadConfig()

	logger, _ := zap.NewProduction()

	defer logger.Sync()
	log := logger.Sugar()

	db := mongodb.Run(cfg)
	rdb := redis.Run(cfg)

	app := fiber.New()
	app.Use(cors.New())

	validate := validator.New()
	repo := repository.NewRepository(cfg, db, rdb)
	services := service.NewService(log, repo)
	handlers := handler.NewHandler(log, services, validate)

	handlers.InitRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", cfg.HttpServer.Host, cfg.HttpServer.Port)))
}
