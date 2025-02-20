package service

import (
	"fmt"

	"github.com/kyamisama/inventory-api/models"
	"github.com/kyamisama/inventory-api/repository"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
}

type ItemMemoryService struct {
	repository repository.IItemRepository
}

func NewItemService(repository repository.IItemRepository) IItemService {
	return &ItemMemoryService{repository: repository}
}

func (s *ItemMemoryService) FindAll() (*[]models.Item, error) {
	items, err := s.repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("could not find items: %w", err)
	}
	if len(*items) == 0 {
		return nil, fmt.Errorf("no items found")
	}
	return items, nil
}
