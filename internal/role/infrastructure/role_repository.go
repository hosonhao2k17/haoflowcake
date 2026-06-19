package infrastructure

import (
	"context"
	roleentity "haoflowcake/internal/role/domain/entity"
	"haoflowcake/internal/role/infrastructure/model"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

//func NewRoleRepository(db *gorm.DB) repository.RoleRepository {
//	return &roleRepository{db: db}
//}

func (roleRepo *roleRepository) Create(ctx context.Context, roleEntity *roleentity.RoleEntity) (*roleentity.RoleEntity, error) {
	roleModel := model.ToRoleModel(roleEntity)
	if err := roleRepo.db.WithContext(ctx).Create(roleModel).Error; err != nil {
		return nil, err
	}
	return roleModel.ToRoleEntity(), nil
}
