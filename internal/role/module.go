package role

import (
	"haoflowcake/internal/role/application/usecase"
	"haoflowcake/internal/role/delivery/http"
	"haoflowcake/internal/role/infrastructure"

	"gorm.io/gorm"
)

type Module struct {
	Handler *http.RoleHandler
}

func NewRoleModule(db *gorm.DB) *Module {

	roleRepository := infrastructure.NewRoleRepository(db)
	createRoleUseCase := usecase.NewCreateRoleUseCase(roleRepository)
	roleHandler := http.NewRoleHandler(createRoleUseCase)

	return &Module{
		Handler: roleHandler,
	}
}
