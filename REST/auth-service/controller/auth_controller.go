package controller

import (
	"cisdi-technical-assessment/REST/auth-service/model/dto"
	"cisdi-technical-assessment/REST/auth-service/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status:  false,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.authService.Login(loginRequest)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status:  false,
			Message: "Login failed",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  true,
		Message: "Login successful",
		Data:    response,
	})
}

func (c *AuthController) ValidateToken(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status:  false,
			Message: "Authorization header required",
		})
		return
	}

	// Extract the token from the Authorization header
	// Format: "Bearer {token}"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status:  false,
			Message: "Invalid authorization format",
		})
		return
	}

	token := parts[1]
	userID, err := c.authService.ValidateToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status:  false,
			Message: "Invalid token",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  true,
		Message: "Token is valid",
		Data: map[string]interface{}{
			"user_id": userID,
		},
	})
}
