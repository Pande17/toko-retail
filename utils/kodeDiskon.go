package utils

import (
	"projek/toko-retail/model"
	config "projek/toko-retail/repository/config"
	"time"
)

func CreateKodeDiskon(data model.Diskon) (model.Diskon, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Mysql.DB)

	return data, err
}

func GetDiskon() ([]model.Diskon, error) {
	var Diskon model.Diskon
	return Diskon.GetAll(config.Mysql.DB)
}

func GetDiskonByCode(s string) (model.Diskon, error) {
	diskon := model.Diskon{
		KodeDiskon: s,
	}

	return diskon.GetByCode(config.Mysql.DB)
}

func GetDiskonByID(id uint) (model.Diskon, error) {
	Diskon := model.Diskon{
		ID: id,
	}
	return Diskon.GetByID(config.Mysql.DB)
}

func UpdateDiskon(id uint, updatedDiskon model.Diskon) (model.Diskon, error) {
	existingDiskon := model.Diskon{ID: id}
	if err := config.Mysql.DB.First(&existingDiskon).Error; err != nil {
		return model.Diskon{}, err
	}

	existingDiskon.Amount = updatedDiskon.Amount
	existingDiskon.Type = updatedDiskon.Type
	existingDiskon.UpdatedAt = time.Now()

	if err := config.Mysql.DB.Save(&existingDiskon).Error; err != nil {
		return model.Diskon{}, err
	}

	return existingDiskon, nil
}

func DeleteKode(id uint64) error {
	Kode := model.Diskon{
		ID: uint(id),
	}
	return Kode.Delete(config.Mysql.DB)
}