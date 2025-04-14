package config

import (
	"cisdi-technical-assessment/REST/auth-service/model/dto"
	"os"
	"strconv"
	"time"
)

func LoadConfig() (*dto.Config, error) {
	var cfg dto.Config

	// Database config
	cfg.Database.Host = getEnv("DB_HOST", "localhost")
	cfg.Database.Port = getEnvAsInt("DB_PORT", 5432)
	cfg.Database.Username = getEnv("DB_USER", "postgres")
	cfg.Database.Password = getEnv("DB_PASSWORD", "postgres")
	cfg.Database.Name = getEnv("DB_NAME", "auth_service")
	cfg.Database.SSLMode = getEnv("DB_SSLMODE", "disable")

	// Server config
	cfg.Server.Host = getEnv("SERVER_HOST", "0.0.0.0")
	cfg.Server.Port = getEnvAsInt("SERVER_PORT", 8080)

	// JWT config
	cfg.JWT.Secret = getEnv("JWT_SECRET", "your-256-bit-secret")
	cfg.JWT.ExpiresIn = time.Duration(getEnvAsInt("JWT_EXPIRES_IN_HOURS", 24)) * time.Hour

	return &cfg, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}
