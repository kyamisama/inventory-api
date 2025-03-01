package service

import (
	"fmt"

	"github.com/kyamisama/inventory-api/dto"
	"github.com/kyamisama/inventory-api/models"
	"github.com/kyamisama/inventory-api/repository"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	CreateItem(dto *dto.CreateItemDto) (*models.Item, error)
	UpdateItem(itemId uint, dto *dto.UpdateItemDto) (*models.Item, error)
	DeleteItem(itemId uint) error
}

type ItemService struct {
	repository repository.IItemRepository
}

func NewItemMemoryService(repository repository.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) FindAll() (*[]models.Item, error) {
	items, err := s.repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("could not find items: %w", err)
	}
	if len(*items) == 0 {
		return nil, fmt.Errorf("no items found")
	}
	return items, nil
}

func (s *ItemService) FindById(itemId uint) (*models.Item, error) {
	item, err := s.repository.FindById(itemId)
	if err != nil {
		return nil, fmt.Errorf("could not find item with ID %d: %w", itemId, err)
	}
	return item, nil
}

func (s *ItemService) CreateItem(dto *dto.CreateItemDto) (*models.Item, error) {
	item := models.Item{
		Name:        dto.Name,
		Description: dto.Description,
		Quantity:    dto.Quantity,
		CreatedBy:   dto.CreatedBy,
	}
	return s.repository.CreateItem(item)
}

func (s *ItemService) UpdateItem(itemId uint, dto *dto.UpdateItemDto) (*models.Item, error) {
	targetItem, err := s.repository.FindById(itemId)
	if err != nil {
		return nil, err
	}
	// 異なるフィールドだけを更新
	if targetItem.Name != dto.Name {
		targetItem.Name = dto.Name
	}
	if targetItem.Description != dto.Description {
		targetItem.Description = dto.Description
	}
	if targetItem.Quantity != dto.Quantity {
		targetItem.Quantity = dto.Quantity
	}
	if targetItem.UpdatedBy != dto.UpdatedBy {
		targetItem.UpdatedBy = dto.UpdatedBy
	}

	return s.repository.UpdateItem(targetItem)
}

func (s *ItemService) DeleteItem(itemId uint) error {
	return s.repository.DeleteItem(itemId)
}
