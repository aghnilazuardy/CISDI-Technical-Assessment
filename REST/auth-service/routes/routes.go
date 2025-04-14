package routes

import (
	"cisdi-technical-assessment/REST/auth-service/controller"
	"cisdi-technical-assessment/REST/auth-service/model/dto"
	"cisdi-technical-assessment/REST/auth-service/repository"
	"cisdi-technical-assessment/REST/auth-service/service"
	"cisdi-technical-assessment/REST/auth-service/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRouter(cfg *dto.Config) *gin.Engine {
	r := gin.Default()

	// Connect to the database
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
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	// Setup repositories
	userRepo := repository.NewUserRepository(db)

	// Setup utilities
	jwtUtil := utils.NewJWTUtil(cfg.JWT.Secret, cfg.JWT.ExpiresIn)

	// Setup services
	authService := service.NewAuthService(userRepo, jwtUtil)

	// Setup controllers
	authController := controller.NewAuthController(authService)

	// Define routes
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.POST("/login", authController.Login)
				auth.POST("/validate", authController.ValidateToken)
			}
		}
	}

	return r
}
