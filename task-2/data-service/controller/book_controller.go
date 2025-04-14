package controller

import (
	"cisdi-technical-assessment/REST/auth-service/validator"
	"cisdi-technical-assessment/REST/data-service/model/dto"
	"cisdi-technical-assessment/REST/data-service/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	service service.BookService
}

func NewBookController(service service.BookService) *BookController {
	return &BookController{service}
}

func (c *BookController) GetAllBooks(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	// Ensure valid pagination parameters
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	response, err := c.service.GetAllBooks(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Status:  false,
			Message: "Failed to fetch books",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  true,
		Message: "Books fetched successfully",
		Data:    response,
	})
}

func (c *BookController) GetBookByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status:  false,
			Message: "Invalid book ID",
			Error:   err.Error(),
		})
		return
	}

	book, err := c.service.GetBookByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status:  false,
			Message: "Book not found",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  true,
		Message: "Book fetched successfully",
		Data:    book,
	})
}

func (c *BookController) CreateBook(ctx *gin.Context) {
	var request dto.CreateBookRequest

	if !validator.ValidateRequest(ctx, &request) {
		return
	}

	book, err := c.service.CreateBook(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status:  false,
			Message: "Failed to create book",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Status:  true,
		Message: "Book created successfully",
		Data:    book,
	})
}
