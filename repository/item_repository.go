package repository

import (
	"errors"
	"fmt"

	"github.com/kyamisama/inventory-api/models"
	"gorm.io/gorm"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	CreateItem(newItem models.Item) (*models.Item, error)
	UpdateItem(updateItem *models.Item) (*models.Item, error)
	DeleteItem(itemId uint) error
}

type ItemMemoryRepository struct {
	items []models.Item
}

func NewItemMemoryRepository(item []models.Item) IItemRepository {
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

func (r *ItemMemoryRepository) CreateItem(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(r.items) + 1)
	r.items = append(r.items, newItem)
	return &newItem, nil
}

func (r *ItemMemoryRepository) UpdateItem(updateItem *models.Item) (*models.Item, error) {
	for i, v := range r.items {
		if v.ID == updateItem.ID {
			r.items[i] = *updateItem
			return &r.items[i], nil
		}
	}
	return nil, errors.New("Unexpected error")
}

func (r *ItemMemoryRepository) DeleteItem(itemId uint) error {
	for i, v := range r.items {
		if v.ID == itemId {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("item ID not found")
}

type ItemRepository struct {
	db *gorm.DB
}

func (r *ItemRepository) CreateItem(newItem models.Item) (*models.Item, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

func (r *ItemRepository) DeleteItem(itemId uint) error {
	result := r.db.Unscoped().Delete(&models.Item{}, itemId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ItemRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := r.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}

func (r *ItemRepository) FindById(itemId uint) (*models.Item, error) {
	var item models.Item
	result := r.db.First(&item, "id = ?", itemId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("Item not found")
		}
		return nil, result.Error
	}
	return &item, nil
}

func (r *ItemRepository) UpdateItem(updateItem *models.Item) (*models.Item, error) {
	result := r.db.Model(&models.Item{}).
		Where("id = ?", updateItem.ID).
		Omit("id", "created_at").
		Updates(updateItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return updateItem, nil
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{db: db}
}
