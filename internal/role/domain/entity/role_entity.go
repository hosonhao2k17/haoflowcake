package roleentity

import (
	"haoflowcake/common/entity"

	"github.com/google/uuid"
)

type RoleEntity struct {
	entity.BaseEntity
	ID          uuid.UUID
	Name        string
	Key         string
	Description *string
}
