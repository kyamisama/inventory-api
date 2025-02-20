package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyamisama/inventory-api/service"
)

type IItemController interface {
	FindAll(ctx *gin.Context)
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
