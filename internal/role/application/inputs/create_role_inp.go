package inputs

import roleentity "haoflowcake/internal/role/domain/entity"

type CreateRoleInput struct {
	Name        string
	Key         string
	Description *string
}

func (createRoleInp *CreateRoleInput) ToEntity() *roleentity.RoleEntity {
	return &roleentity.RoleEntity{
		Name:        createRoleInp.Name,
		Key:         createRoleInp.Key,
		Description: createRoleInp.Description,
	}
}
