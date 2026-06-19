package dto

import "haoflowcake/internal/role/application/inputs"

type CreateRoleDto struct {
	Name        string  `json:"name"`
	Key         string  `json:"key"`
	Description *string `json:"description"`
}

func (dto *CreateRoleDto) ToCreateRoleInput() *inputs.CreateRoleInput {
	return &inputs.CreateRoleInput{
		Name:        dto.Name,
		Key:         dto.Key,
		Description: dto.Description,
	}
}
