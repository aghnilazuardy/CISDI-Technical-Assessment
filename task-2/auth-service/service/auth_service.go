package service

import (
	"cisdi-technical-assessment/REST/auth-service/model/dto"
	"cisdi-technical-assessment/REST/auth-service/repository"
	"cisdi-technical-assessment/REST/auth-service/utils"
	"errors"

	"gorm.io/gorm"
)

type AuthService interface {
	Login(loginRequest dto.LoginRequest) (*dto.LoginResponse, error)
	ValidateToken(token string) (uint, error)
}

type authService struct {
	userRepo repository.UserRepository
	jwtUtil  utils.JWTUtil
}

func NewAuthService(userRepo repository.UserRepository, jwtUtil utils.JWTUtil) AuthService {
	return &authService{
		userRepo: userRepo,
		jwtUtil:  jwtUtil,
	}
}

func (s *authService) Login(loginRequest dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepo.FindByUsername(loginRequest.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid username or password")
		}
		return nil, err
	}

	if !utils.CheckPasswordHash(loginRequest.Password, user.Password) {
		return nil, errors.New("invalid username or password")
	}

	token, err := s.jwtUtil.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{Token: token}, nil
}

func (s *authService) ValidateToken(token string) (uint, error) {
	parsedToken, err := s.jwtUtil.ValidateToken(token)
	if err != nil {
		return 0, err
	}

	userID, err := s.jwtUtil.ExtractUserID(parsedToken)
	if err != nil {
		return 0, err
	}

	// Verify that the user exists
	_, err = s.userRepo.FindByID(userID)
	if err != nil {
		return 0, errors.New("user not found")
	}

	return userID, nil
}
