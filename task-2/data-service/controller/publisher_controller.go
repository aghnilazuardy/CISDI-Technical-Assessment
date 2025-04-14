package controller

import (
	"cisdi-technical-assessment/REST/auth-service/validator"
	"cisdi-technical-assessment/REST/data-service/model/dto"
	"cisdi-technical-assessment/REST/data-service/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PublisherController struct {
	service service.PublisherService
}

func NewPublisherController(service service.PublisherService) *PublisherController {
	return &PublisherController{service}
}

func (c *PublisherController) GetAllPublishers(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	// Ensure valid pagination parameters
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	response, err := c.service.GetAllPublishers(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Status:  false,
			Message: "Failed to fetch publishers",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  true,
		Message: "Publishers fetched successfully",
		Data:    response,
	})
}

func (c *PublisherController) GetPublisherByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status:  false,
			Message: "Invalid publisher ID",
			Error:   err.Error(),
		})
		return
	}

	publisher, err := c.service.GetPublisherByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status:  false,
			Message: "Publisher not found",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  true,
		Message: "Publisher fetched successfully",
		Data:    publisher,
	})
}

func (c *PublisherController) CreatePublisher(ctx *gin.Context) {
	var request dto.CreatePublisherRequest

	if !validator.ValidateRequest(ctx, &request) {
		return
	}

	publisher, err := c.service.CreatePublisher(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status:  false,
			Message: "Failed to create publisher",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Status:  true,
		Message: "Publisher created successfully",
		Data:    publisher,
	})
}

func (c *PublisherController) UpdatePublisher(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status:  false,
			Message: "Invalid publisher ID",
			Error:   err.Error(),
		})
		return
	}

	var request dto.UpdatePublisherRequest
	if !validator.ValidateRequest(ctx, &request) {
		return
	}

	publisher, err := c.service.UpdatePublisher(uint(id), request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status:  false,
			Message: "Failed to update publisher",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  true,
		Message: "Publisher updated successfully",
		Data:    publisher,
	})
}

func (c *PublisherController) DeletePublisher(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status:  false,
			Message: "Invalid publisher ID",
			Error:   err.Error(),
		})
		return
	}

	err = c.service.DeletePublisher(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status:  false,
			Message: "Failed to delete publisher",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  true,
		Message: "Publisher deleted successfully",
	})
}
