package item

import (
	"content-editor/internal/domain/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Begin() (*gorm.DB, error) {
	return r.begin()
}

func (r *Repository) begin() (*gorm.DB, error) {
	tx := r.db.Begin()
	return tx, tx.Error
}

func (r *Repository) Create(item *models.Item) (*models.Item, error) {
	return r.CreateTx(item, r.db)
}

func (r *Repository) CreateTx(item *models.Item, tx *gorm.DB) (*models.Item, error) {
	if err := tx.Create(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}
