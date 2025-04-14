package validator

import (
	"cisdi-technical-assessment/REST/auth-service/model/dto"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateRequest(c *gin.Context, request interface{}) bool {
	if err := c.ShouldBindJSON(&request); err != nil {
		var validationErrors []dto.ValidationError

		if verr, ok := err.(validator.ValidationErrors); ok {
			for _, f := range verr {
				var message string
				field := strings.ToLower(f.Field())

				switch f.Tag() {
				case "required":
					message = field + " is required"
				case "email":
					message = field + " should be a valid email"
				case "min":
					message = field + " should have at least " + f.Param() + " characters"
				case "max":
					message = field + " should have at most " + f.Param() + " characters"
				default:
					message = field + " is not valid"
				}

				validationErrors = append(validationErrors, dto.ValidationError{
					Field:   field,
					Message: message,
				})
			}
		} else {
			c.JSON(http.StatusBadRequest, dto.Response{
				Status:  false,
				Message: "Invalid request",
				Error:   err.Error(),
			})
			return false
		}

		c.JSON(http.StatusBadRequest, dto.Response{
			Status:  false,
			Message: "Validation failed",
			Error:   validationErrors,
		})
		return false
	}
	return true
}
