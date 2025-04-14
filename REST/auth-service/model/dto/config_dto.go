package dto

import "time"

type Config struct {
	Database struct {
		Host     string
		Port     int
		Username string
		Password string
		Name     string
		SSLMode  string
	}
	Server struct {
		Host string
		Port int
	}
	JWT struct {
		Secret    string
		ExpiresIn time.Duration
	}
}
