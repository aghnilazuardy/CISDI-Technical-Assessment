package dto

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
	Auth struct {
		ServiceURL string
	}
}
