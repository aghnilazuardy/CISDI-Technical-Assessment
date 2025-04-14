package controller

import (
	"cisdi-technical-assessment/REST/auth-service/validator"
	"cisdi-technical-assessment/REST/data-service/model/dto"
	"cisdi-technical-assessment/REST/data-service/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	service service.AuthorService
}

func NewAuthorController(service service.AuthorService) *AuthorController {
	return &AuthorController{service}
}

func (c *AuthorController) GetAllAuthors(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	// Ensure valid pagination parameters
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	response, err := c.service.GetAllAuthors(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Status:  false,
			Message: "Failed to fetch authors",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  true,
		Message: "Authors fetched successfully",
		Data:    response,
	})
}

func (c *AuthorController) GetAuthorByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status:  false,
			Message: "Invalid author ID",
			Error:   err.Error(),
		})
		return
	}

	author, err := c.service.GetAuthorByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status:  false,
			Message: "Author not found",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  true,
		Message: "Author fetched successfully",
		Data:    author,
	})
}

func (c *AuthorController) CreateAuthor(ctx *gin.Context) {
	var request dto.CreateAuthorRequest

	if !validator.ValidateRequest(ctx, &request) {
		return
	}

	author, err := c.service.CreateAuthor(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status:  false,
			Message: "Failed to create author",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Status:  true,
		Message: "Author created successfully",
		Data:    author,
	})
}

func (c *AuthorController) UpdateAuthor(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status:  false,
			Message: "Invalid author ID",
			Error:   err.Error(),
		})
		return
	}

	var request dto.UpdateAuthorRequest
	if !validator.ValidateRequest(ctx, &request) {
		return
	}

	author, err := c.service.UpdateAuthor(uint(id), request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status:  false,
			Message: "Failed to update author",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  true,
		Message: "Author updated successfully",
		Data:    author,
	})
}

func (c *AuthorController) DeleteAuthor(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status:  false,
			Message: "Invalid author ID",
			Error:   err.Error(),
		})
		return
	}

	err = c.service.DeleteAuthor(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status:  false,
			Message: "Failed to delete author",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  true,
		Message: "Author deleted successfully",
	})
}
