package migrations

import (
	"cisdi-technical-assessment/REST/data-service/model"
	"cisdi-technical-assessment/REST/data-service/model/dto"
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
	err = db.AutoMigrate(&model.Book{}, &model.Author{}, &model.Publisher{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
