package models

import "time"

type Manufacturer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	OfficialName string `json:"officialName"`
	Address      string `json:"address"`

	UpdatedAt *time.Time `json:"updatedAt" gorm:"type:timestamp"`
	CreatedAt time.Time  `json:"createdAt" gorm:"type:timestamp"`
}
