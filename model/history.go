package model

import "gorm.io/gorm"

type Histori struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	ID_barang  uint   `json:"id_barang"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Keterangan string `json:"keterangan"`
	Model
}

type HistoriASK struct {
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Keterangan string `json:"keterangan"`
}

type HistoriASKM struct {
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Keterangan string `json:"keterangan"`
	Model
}

func (hs *Histori) Create(db *gorm.DB) error {
	err := db.Model(Histori{}).Create(&hs).Error
	if err != nil {
		return err
	}
	return nil
}

func (hs *Histori) GetIDBarang (db *gorm.DB) ([]Histori, error) {
	res := []Histori{}
	err := db.Model(Histori{}).Where("id_barang = ?", hs.ID_barang).Find(&res).Error
	if err!= nil {
        return []Histori{}, err
    }
	return res, err
}