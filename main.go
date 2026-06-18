package main

import (
	"fmt"
	"haoflowcake/config"

	"github.com/gofiber/fiber/v3"
)

func main() {
	cf := config.LoadConfig()
	app := fiber.New()
	fmt.Println(cf.App)
	app.Listen(cf.App.GetPort())
}
