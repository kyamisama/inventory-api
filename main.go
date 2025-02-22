package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kyamisama/inventory-api/controller"
	"github.com/kyamisama/inventory-api/models"
	"github.com/kyamisama/inventory-api/repository"
	"github.com/kyamisama/inventory-api/service"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "item1", Description: "item1 desc", Quantity: 1, CreatedBy: "yamada"},
		{ID: 2, Name: "item2", Description: "item2 desc", Quantity: 2, CreatedBy: "tanaka"},
	}
	itemRepository := repository.NewItemRepository(items)
	itemService := service.NewItemService(itemRepository)
	itemController := controller.NewItemController(itemService)
	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.POST("", itemController.CreateItem)
	r.Run("localhost:8000") // 0.0.0.0:8080 でサーバーを立てます。
}
