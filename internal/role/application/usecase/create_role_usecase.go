package usecase

import (
	"context"
	"haoflowcake/internal/role/application/inputs"
	roleentity "haoflowcake/internal/role/domain/entity"
	"haoflowcake/internal/role/domain/repository"
)

type CreateRoleUseCase interface {
	Execute(ctx context.Context, createRoleInp *inputs.CreateRoleInput) (*roleentity.RoleEntity, error)
}

type createRoleUseCase struct {
	roleRepository repository.RoleRepository
}

func NewCreateRoleUseCase(repository repository.RoleRepository) CreateRoleUseCase {
	return &createRoleUseCase{
		roleRepository: repository,
	}
}

func (createUseCase *createRoleUseCase) Execute(ctx context.Context, createRoleInp *inputs.CreateRoleInput) (*roleentity.RoleEntity, error) {
	return createUseCase.roleRepository.Create(ctx, createRoleInp.ToEntity())
}
