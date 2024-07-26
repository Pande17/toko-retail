package utils

import (
	"projek/toko-retail/model"
	repository "projek/toko-retail/repository/config"
	"time"
)

// function untuk input kode diskon
func CreateKodeDiskon(data model.Diskon) (model.Diskon, error) {
		data.CreatedAt = time.Now()
		data.UpdatedAt = time.Now()
		err := data.Create(repository.Mysql.DB)

		return data, err
}

// function untuk menampilkan data kode diskon
func GetDiskon() ([]model.Diskon, error) {
		var Diskon model.Diskon
		return Diskon.GetAll(repository.Mysql.DB)
}

// function untuk menampilkan data kode diskon berdasarkan ID-nya
func GetDiskonByID(id uint) (model.Diskon, error) {
		Diskon := model.Diskon{
				ID: id,
		}
		return Diskon.GetByID(repository.Mysql.DB)
}

// function untuk edit/update data kode diskon
func UpdateDiskon(id uint, updateDiskon model.Diskon) (model.Diskon, error) {
		existingDiskon := model.Diskon{ID: id}
		if err := repository.Mysql.DB.First(&existingDiskon).Error; err != nil {
				return model.Diskon{}, err
		}

		existingDiskon.Amount 		= updateDiskon.Amount
		existingDiskon.Type 		= updateDiskon.Type
		existingDiskon.UpdatedAt	= time.Now()

		if err := repository.Mysql.DB.Save(&existingDiskon).Error; err != nil {
				return model.Diskon{}, err
		}

		return existingDiskon, nil
}

// function untuk menghapus kode diskon
func DeleteDiskon(id uint64) error {
		Kode := model.Diskon{
				ID: uint(id),
		}
		return Kode.Delete(repository.Mysql.DB)
}