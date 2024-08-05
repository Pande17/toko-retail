package modelfunc

import (
	"projek/toko-retail/model"

	"gorm.io/gorm"
)

type Histori struct {
	model.Histori
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