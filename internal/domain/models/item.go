package models

import "time"

type Item struct {
	ID             int        `json:"id"`
	Code           string     `json:"code"`
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	Type           int        `json:"type"`
	ManufacturerID int        `json:"manufactirerID"`
	UpdatedAt      *time.Time `json:"updatedAt" gorm:"type:timestamp"`
	CreatedAt      time.Time  `json:"createdAt" gorm:"type:timestamp"`

	Manufacturer *Manufacturer `json:"-" gorm:"->;foreign key:ManufacturerID"`
}
