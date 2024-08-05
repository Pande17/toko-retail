package modelfunc

import (
	"projek/toko-retail/model"

	"gorm.io/gorm"
)

type Penjualan struct {
	model.Penjualan
}

func (pj *Penjualan) CreatePenjualan(db *gorm.DB) error {
	err := db.Create(&pj).Error
	if err != nil {
		return err
	}

	return nil
}

func (pj *Penjualan) GetAll(db *gorm.DB) ([]Penjualan, error) {
	res := []Penjualan{}

	err := db.Model(Penjualan{}).Find(&res).Error
	if err != nil {
		return []Penjualan{}, err
	}

	return res, nil
}

func (pj *Penjualan) GetPByID(db *gorm.DB) (Penjualan, error) {
	res := Penjualan{}

	err := db.Model(Penjualan{}).Where("id = ?", pj.ID).Take(&res).Error
	if err != nil {
		return Penjualan{}, err
	}

	return res, nil
}

func (pj *Penjualan) Update(db *gorm.DB) error {
	err := db.Model(Penjualan{}).Where("id = ?", pj.ID).Updates(&pj).Error
	if err != nil {
		return err
	}
	return nil
}