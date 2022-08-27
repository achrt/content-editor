package models

type Manufacturer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	OfficialName string `json:"officialName"`
	Address      string `json:"address"`
}
