package main

import (
	"haoflowcake/config"
	"haoflowcake/database"
	"haoflowcake/logger"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func main() {
	cf := config.LoadConfig()
	app := fiber.New()
	lg := logger.NewLogger(cf)
	lg.Info("init haoflowcake", zap.String("port", cf.App.GetPort()))

	//CONNECT TO DATABASE
	database.Connect(cf, lg)
	if err := app.Listen(cf.App.GetPort()); err != nil {
		lg.Fatal("fail when init app", zap.Error(err))
	}
}
