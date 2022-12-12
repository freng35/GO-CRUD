package database

import (
	"app/pkg/config"
	"app/pkg/model"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func createTables(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{})

	return err
}

func CreateDBConnection(cfg *config.Config, logger *zap.SugaredLogger) error {
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.DBAddress), &gorm.Config{})
	if err != nil {
		logger.Fatalw("Failed on connection to db", "err", err)
	}

	if err = createTables(DB); err != nil {
		logger.Fatalw("Failed on tables migration", "err", err)
	}

	return err
}
