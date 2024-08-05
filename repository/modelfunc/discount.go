package modelfunc

import (
	"projek/toko-retail/model"

	"gorm.io/gorm"
)

type Diskon struct {
	model.Diskon
}

func (kd *Diskon) CreateDiskon(db *gorm.DB) error {
	err := db.Model(Diskon{}).Create(&kd).Error
	if err != nil {
		return err
	}

	return nil
}

func (kd *Diskon) GetAll(db *gorm.DB) ([]Diskon, error) {
	res := []Diskon{}

	err := db.
		Model(Diskon{}).
		Find(&res).
		Error

	if err != nil {
		return []Diskon{}, err
	}

	return res, nil
}

func (kd *Diskon) GetByCode(db *gorm.DB) (Diskon, error) {
	res := Diskon{}

	err := db.
		Where("kode_diskon = ?", kd.KodeDiskon).
		First(&res).
		Error

	if err != nil {
		return Diskon{}, err
	}

	return res, nil
}

func (kd *Diskon) GetByID(db *gorm.DB) (Diskon, error) {
	res := Diskon{}

	err := db.
		Model(Diskon{}).
		Where("id = ?", kd.ID).
		Take(&res).
		Error

	if err != nil {
		return Diskon{}, err
	}

	return res, nil
}

func (kd *Diskon) Delete(db *gorm.DB) error {
	err := db.
		Model(&Diskon{}).
		Where("id = ?", kd.ID).
		Delete(&kd).
		Error

	if err != nil {
		return err
	}

	return nil
}