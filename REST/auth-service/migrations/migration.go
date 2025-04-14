package migrations

import (
	"cisdi-technical-assessment/REST/auth-service/model"
	"cisdi-technical-assessment/REST/auth-service/model/dto"
	"cisdi-technical-assessment/REST/auth-service/utils"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RunMigrations(cfg *dto.Config) error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Run migrations
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	// Seed admin user if not exists
	var count int64
	db.Model(&model.User{}).Count(&count)
	if count == 0 {
		hashedPassword, err := utils.HashPassword("admin123")
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}

		adminUser := &model.User{
			Username: "admin",
			Email:    "admin@example.com",
			Password: hashedPassword,
		}

		if err := db.Create(adminUser).Error; err != nil {
			return fmt.Errorf("failed to create admin user: %w", err)
		}
	}

	return nil
}
