package models

import "time"

type Catalogue struct {
	ID        int        `json:"id"`
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"type:timestamp"`
	CreatedAt time.Time  `json:"createdAt" gorm:"type:timestamp"`

	Items []*Item `json:"items" gorm:"many2many:catalogue_items;"`
}
