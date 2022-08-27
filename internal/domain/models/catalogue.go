package models

type Catalogue struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`

	Items []*Item `json:"items" gorm:"many2many:catalogue_items;"`
}
