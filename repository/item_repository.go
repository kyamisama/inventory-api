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
	UpdateItem(dto *dto.UpdateItemDto) (*models.Item, error)
	DeleteItem(itemId uint) (*models.Item, error)
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

func (r *ItemMemoryRepository) UpdateItem(dto *dto.UpdateItemDto) (*models.Item, error) {
	for idx, item := range r.items {
		if item.ID != dto.ID {
			continue // IDが一致しない場合、次のアイテムにスキップ
		}

		// 異なるフィールドだけを更新
		if item.Name != dto.Name {
			item.Name = dto.Name
		}
		if item.Description != dto.Description {
			item.Description = dto.Description
		}
		if item.Quantity != dto.Quantity {
			item.Quantity = dto.Quantity
		}
		if item.UpdatedBy != dto.UpdatedBy {
			item.UpdatedBy = dto.UpdatedBy
		}

		// 更新したアイテムを返す
		r.items[idx] = item
		return &item, nil
	}
	return nil, errors.New("item not found")
}

func (r *ItemMemoryRepository) DeleteItem(itemId uint) (*models.Item, error) {
	for i, v := range r.items {
		if v.ID == itemId {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return &v, nil
		}
	}
	return nil, fmt.Errorf("item ID not found")
}
