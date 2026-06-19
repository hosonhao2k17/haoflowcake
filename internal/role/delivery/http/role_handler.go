package http

import (
	"haoflowcake/common"
	"haoflowcake/internal/role/application/usecase"
	"haoflowcake/internal/role/delivery/dto"

	"github.com/gofiber/fiber/v3"
)

type RoleHandler struct {
	createRoleUseCase usecase.CreateRoleUseCase
}

func NewRoleHandler(
	createRoleUseCase usecase.CreateRoleUseCase,
) *RoleHandler {

	return &RoleHandler{
		createRoleUseCase: createRoleUseCase,
	}
}

func (handler *RoleHandler) Create(c fiber.Ctx) error {
	createRoleDto := new(dto.CreateRoleDto)
	if err := c.Bind().Body(createRoleDto); err != nil {
		return err
	}
	data, err := handler.createRoleUseCase.Execute(c.Context(), createRoleDto.ToCreateRoleInput())
	if err != nil {
		return err
	}
	return common.ResponseCreated(c, data, "created role success")
}
