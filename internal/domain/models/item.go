package models

type Item struct {
	ID             int    `json:"id"`
	Code           string `json:"code"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Type           int    `json:"type"`
	ManufacturerID int    `json:"manufactirerID"`

	Manufacturer *Manufacturer `json:"-" gorm:"->;foreignkey:ManufacturerID"`
}
