package modelfunc

import (
	"projek/toko-retail/model"

	"gorm.io/gorm"
)

type Barang struct {
	model.Barang
}

func (br *Barang) Create(db *gorm.DB) error {
	err := db.Model(model.Barang{}).Create(&br).Error
	if err != nil {
		return err
	}
	return nil
}

func (br *Barang) GetAll(db *gorm.DB) ([]model.Barang, error) {
	res := []model.Barang{}
	err := db.Model(model.Barang{}).Find(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (br *Barang) GetByID(db *gorm.DB) (model.Barang, error) {
	res := model.Barang{}
	err := db.Model(model.Barang{}).Where("id = ?", br.ID).Find(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (br *Barang) Update(db *gorm.DB) error {
	err := db.Model(model.Barang{}).Where("id = ?", br.ID).Updates(&br).Error
	if err != nil {
		return err
	}
	return nil
}

func (br *Barang) Delete(db *gorm.DB) error {
	err := db.Where("id = ?", br.ID).Delete(&model.Barang{}).Error
	if err != nil {
		return err
	}
	return nil
}