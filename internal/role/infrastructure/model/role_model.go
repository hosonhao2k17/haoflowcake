package model

import (
	"haoflowcake/common/entity"
	commonmodel "haoflowcake/database/model"
	"haoflowcake/internal/role/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleModel struct {
	commonmodel.BaseModel
	ID          uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string         `json:"name" gorm:"not null"`
	Key         string         `json:"key" gorm:"unique,not null"`
	Description *string        `json:"description"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func (RoleModel) TableName() string {
	return "roles"
}

func (roleModel *RoleModel) ToRoleEntity() *roleentity.RoleEntity {
	return &roleentity.RoleEntity{
		ID:          roleModel.ID,
		Name:        roleModel.Name,
		Key:         roleModel.Key,
		Description: roleModel.Description,
		BaseEntity: entity.BaseEntity{
			CreatedAt: roleModel.CreatedAt,
			UpdateAt:  roleModel.UpdatedAt,
		},
	}
}

func ToRoleModel(roleEntity *roleentity.RoleEntity) *RoleModel {
	return &RoleModel{
		ID:          roleEntity.ID,
		Name:        roleEntity.Name,
		Key:         roleEntity.Key,
		Description: roleEntity.Description,
		BaseModel: commonmodel.BaseModel{
			CreatedAt: roleEntity.CreatedAt,
			UpdatedAt: roleEntity.UpdateAt,
		},
	}
}
