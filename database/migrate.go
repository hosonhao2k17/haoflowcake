package database

import (
	"haoflowcake/internal/role/infrastructure/model"
	"haoflowcake/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB, logger *logger.Logger) {
	if err := db.AutoMigrate(&model.RoleModel{}); err != nil {
		logger.Fatal("auto migrate fail", zap.Error(err))
	}

	logger.Info("auto migrate success")
}
