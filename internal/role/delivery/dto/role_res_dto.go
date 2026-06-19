package dto

import (
	roleentity "haoflowcake/internal/role/domain/entity"
	"time"

	"github.com/google/uuid"
)

type RoleResponseDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Key         string    `json:"key"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewRoleResponse(entity *roleentity.RoleEntity) *RoleResponseDto {
	return &RoleResponseDto{
		ID:          entity.ID,
		Name:        entity.Name,
		Key:         entity.Key,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdateAt,
	}
}
