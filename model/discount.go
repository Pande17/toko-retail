package model

type Diskon struct {
	ID         uint    `gorm:"primarykey" json:"id"`
	KodeDiskon string  `json:"kode_diskon"`
	Amount     float64 `json:"amount"`
	Type       string  `json:"type"`
	Model
}

