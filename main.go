package main

import (
	"github.com/alifiroozi80/duckdb"
	"github.com/gin-gonic/gin"
	"github.com/kyamisama/inventory-api/controller"

	// "github.com/kyamisama/inventory-api/models"
	"github.com/kyamisama/inventory-api/repository"
	"github.com/kyamisama/inventory-api/service"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(duckdb.Open("inventory-api.duckdb"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// items := []models.Item{
	// 	{ID: 1, Name: "item1", Description: "item1 desc", Quantity: 1, CreatedBy: "yamada"},
	// 	{ID: 2, Name: "item2", Description: "item2 desc", Quantity: 2, CreatedBy: "tanaka"},
	// }
	// itemRepository := repository.NewItemMemoryRepository(items)
	itemRepository := repository.NewItemRepository(db)
	itemService := service.NewItemMemoryService(itemRepository)
	itemController := controller.NewItemMemoryController(itemService)
	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.POST("", itemController.CreateItem)
	r.PUT("/items/:id", itemController.UpdateItem)
	r.DELETE("/items/:id", itemController.DeleteItem)
	r.Run("localhost:8000") // 0.0.0.0:8080 でサーバーを立てます。
}
