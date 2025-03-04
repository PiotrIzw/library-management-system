package database

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	bModel "library-management-system/internal/book/model"
	"library-management-system/internal/shared/config"
	uModel "library-management-system/internal/user/model"
)

var DB *gorm.DB

func ConnectDB(logger *zap.Logger) {
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load config", zap.Error(err))
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect database", zap.Error(err))
	}

	db.AutoMigrate(&uModel.User{}, &bModel.Book{})
	DB = db
	logger.Info("Database connection established")
}
