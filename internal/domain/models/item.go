package models

import "time"

type Item struct {
	ID             int        `json:"id"`
	Code           string     `json:"code"`
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	Type           int        `json:"type"`
	ManufacturerID int        `json:"manufactirerID"`
	UpdatedAt      *time.Time `json:"updatedAt" gorm:"type:timestampz"`
	CreatedAt      time.Time  `json:"createdAt" gorm:"type:timestampz"`

	Manufacturer *Manufacturer `json:"-" gorm:"->;foreignkey:ManufacturerID"`
}
