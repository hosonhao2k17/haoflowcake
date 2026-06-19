package repository

import (
	"context"
	roleentity "haoflowcake/internal/role/domain/entity"
)

type RoleRepository interface {
	Create(ctx context.Context, roleEntity *roleentity.RoleEntity) (*roleentity.RoleEntity, error)
}
