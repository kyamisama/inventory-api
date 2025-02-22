package repository

import (
	"errors"
	"fmt"

	"github.com/kyamisama/inventory-api/models"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
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

func (r *ItemMemoryRepository) FindById(itemId uint) (*models.Item, error) {
	for _, item := range r.items {
		if item.ID == itemId {
			return &item, nil
		}
	}
	return nil, errors.New("Item not found")
}
