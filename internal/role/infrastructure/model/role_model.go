package model

import (
	commonmodel "haoflowcake/database/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleModel struct {
	commonmodel.BaseModel
	ID          uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string         `json:"name"`
	Description *string        `json:"description"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func (RoleModel) TableName() string {
	return "roles"
}
