package repository

import (
	"errors"
	"fmt"

	"github.com/kyamisama/inventory-api/dto"
	"github.com/kyamisama/inventory-api/models"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	CreateItem(dto *dto.CreateItemDto) (*models.Item, error)
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

func (r *ItemMemoryRepository) CreateItem(dto *dto.CreateItemDto) (*models.Item, error) {
	item := models.Item{
		Name:        dto.Name,
		Description: dto.Description,
		Quantity:    dto.Quantity,
		CreatedBy:   dto.CreatedBy,
	}
	var maxID uint = 0
	for _, item := range r.items {
		if item.ID > maxID {
			maxID = item.ID
		}
	}
	newID := maxID + 1
	item.ID = newID
	r.items = append(r.items, item)
	return &item, nil
}
