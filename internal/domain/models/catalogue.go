package models

import "time"

type Catalogue struct {
	ID        int        `json:"id"`
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"type:timestampz"`
	CreatedAt time.Time  `json:"createdAt" gorm:"type:timestampz"`

	Items []*Item `json:"items" gorm:"many2many:catalogue_items;"`
}
