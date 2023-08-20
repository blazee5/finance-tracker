package main

import (
	"fmt"
	_ "github.com/blazee5/finance-tracker/docs"
	"github.com/blazee5/finance-tracker/internal/config"
	"github.com/blazee5/finance-tracker/internal/handler"
	"github.com/blazee5/finance-tracker/internal/service"
	storage "github.com/blazee5/finance-tracker/internal/storage/mongodb"
	"github.com/gofiber/fiber/v2"
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

	db, err := storage.Run(cfg)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	userRepo, err := storage.NewUserDAO(db.Db, cfg)
	transactionRepo, err := storage.NewTransactionDAO(db.Db, cfg)
	newStorage := &storage.Storage{Db: db.Db, UserDAO: userRepo, TransactionDAO: transactionRepo}
	services := service.NewService(newStorage)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes(app)

	fmt.Println(service.GenerateHashPassword("password"))

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))

}
