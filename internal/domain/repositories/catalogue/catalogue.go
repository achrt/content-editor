package catalogue

import (
	"gorm.io/gorm"

	"content-editor/internal/domain/models"
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

func (r *Repository) Create(Catalogue *models.Catalogue) (*models.Catalogue, error) {
	return r.CreateTx(Catalogue, r.db)
}

func (r *Repository) CreateTx(Catalogue *models.Catalogue, tx *gorm.DB) (*models.Catalogue, error) {
	if err := tx.Create(&Catalogue).Error; err != nil {
		return nil, err
	}
	return Catalogue, nil
}

// TODO: дописать методы
func (r *Repository) Update(cat *models.Catalogue) error {
	return nil
}

func (r *Repository) Delete(id int) error {
	return nil
}

func (r *Repository) Find(id int) (*models.Catalogue, error) {
	return nil, nil
}

