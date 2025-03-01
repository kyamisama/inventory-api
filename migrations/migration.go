package main

import (
	"github.com/alifiroozi80/duckdb"
	"github.com/kyamisama/inventory-api/models"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(duckdb.Open("inventory-api.duckdb"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	if err := db.AutoMigrate(&models.Item{}); err != nil {
		panic("Failed to migrate database")
	}
}
