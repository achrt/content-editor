package repositories

import (
	"content-editor/internal/domain/models"

	"gorm.io/gorm"
)

type Item interface {
	Begin() (*gorm.DB, error)
	Create(cat *models.Item) (*models.Item, error)
	Update(cat *models.Item) error
	Delete(id int) error
}
