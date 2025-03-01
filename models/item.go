package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Quantity    uint       `json:"quantity"`
	CreatedBy   string     `json:"created_by"`
	UpdatedBy   string     `json:"updated_by"`
	DeletedAt   *time.Time `gorm:"index"`
}

type GormItem struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity" gorm:"not null"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
}
