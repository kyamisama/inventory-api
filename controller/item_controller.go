package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kyamisama/inventory-api/dto"
	"github.com/kyamisama/inventory-api/service"
)

type IItemController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	CreateItem(ctx *gin.Context)
	DeleteItem(ctx *gin.Context)
}

type ItemController struct {
	services service.IItemService
}

func NewItemController(services service.IItemService) IItemController {
	return &ItemController{services: services}
}

func (c *ItemController) FindAll(ctx *gin.Context) {
	items, err := c.services.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

func (c *ItemController) FindById(ctx *gin.Context) {
	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	item, err := c.services.FindById(uint(itemId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": item})
}

func (c *ItemController) CreateItem(ctx *gin.Context) {
	var dto dto.CreateItemDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	item, err := c.services.CreateItem(&dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": item})
}
func (c *ItemController) DeleteItem(ctx *gin.Context) {
	deleteItemId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	item, err := c.services.DeleteItem(uint(deleteItemId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Failed to delete item"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"deleted_item": item})
}
