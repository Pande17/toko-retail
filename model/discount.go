package model

import "gorm.io/gorm"

type Diskon struct {
	ID         uint    `gorm:"primarykey" json:"id"`
	KodeDiskon string  `json:"kode_diskon"`
	Amount     float64 `json:"amount"`
	Type       string  `json:"type"`
	Model
}

func (kd *Diskon) Create(db *gorm.DB) error {
	err := db.
		Model(Diskon{}).
		Create(&kd).
		Error

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