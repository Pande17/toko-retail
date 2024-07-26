package model

import "gorm.io/gorm"

type Penjualan struct {
	ID           uint64  `gorm:"primarykey" json:"id"`
	Kode_invoice string  `json:"kode_invoice"`
	Nama_pembeli string  `json:"nama_pembeli"`
	Subtotal     float64 `json:"subtotal"`
	Kode_diskon  string  `json:"kode_diskon"`
	Diskon       float64 `json:"diskon"`
	Total        float64 `json:"total"`
	Model
	// Diskon
	Created_by string `json:"created_by"`
}

type CreateP struct {
	ID            uint64          `gorm:"primarykey" json:"id"`
	Kode_invoice  string          `json:"kode_invoice"`
	Nama_pembeli  string          `json:"nama_pembeli"`
	Subtotal      float64         `json:"subtotal"`
	Kode_diskon   string          `json:"kode_diskon"`
	Diskon        float64         `json:"diskon"`
	Created_by    string          `json:"created_by"`
	ItemPenjualan []ItemPenjualan `json:"item_penjualan"`
}

func (pj *Penjualan) CreatePenjualan(db *gorm.DB) error {
	err := db.
		Model(Penjualan{}).
		Create(&pj).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (pj *Penjualan) GetAll(db *gorm.DB) ([]Penjualan, error) {
	res := []Penjualan{}

	err := db.
		Model(Penjualan{}).
		Find(&res).
		Error

	if err != nil {
		return []Penjualan{}, err
	}

	return res, nil
}

func (pj *Penjualan) GetPByID(db *gorm.DB) (Penjualan, error) {
	res := Penjualan{}

	err := db.
		Model(Penjualan{}).
		Where("id = ?", pj.ID).
		Take(&res).
		Error

	if err != nil {
		return Penjualan{}, err
	}

	return res, nil
}