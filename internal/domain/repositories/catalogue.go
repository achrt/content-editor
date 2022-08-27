package repositories

import (
	"content-editor/internal/domain/models"

	"gorm.io/gorm"
)

type Catalogue interface {
	Begin() (*gorm.DB, error)
	Find(id int) (*models.Catalogue, error)
	Create(cat *models.Catalogue) (*models.Catalogue, error)
	Update(cat *models.Catalogue) error
	Delete(id int) error
}
