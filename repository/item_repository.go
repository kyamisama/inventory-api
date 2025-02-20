package repository

import (
	"fmt"

	"github.com/kyamisama/inventory-api/models"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
}

type ItemMemoryRepository struct {
	items []models.Item
}

func NewItemRepository(item []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: item}
}

func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	if len(r.items) == 0 {
		return nil, fmt.Errorf("no items found")
	}
	return &r.items, nil
}
