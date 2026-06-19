package repository

import "context"

type RoleRepository interface {
	Create(ctx context.Context)
}
