package router

import (
	"haoflowcake/internal/role"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func RegisterRouter(app *fiber.App, db *gorm.DB) {

	api := app.Group("api")
	v1 := api.Group("v1")

	roleModule := role.NewRoleModule(db)
	roleModule.Handler.RegisterRouter(v1)

}
