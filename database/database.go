package database

import (
	"haoflowcake/config"
	"haoflowcake/logger"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cf *config.Config, lg *logger.Logger) *gorm.DB {
	dsn := cf.Db.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		lg.Fatal("cannot connected to database", zap.Error(err))
	}
	lg.Info("connected to database")
	return db
}
