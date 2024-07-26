package utils

import (
	"projek/toko-retail/model"
	repository "projek/toko-retail/repository/config"
	"time"
)

// function untuk memasukkan data penjualan ke DB
func InsertPenjualanData(data model.Penjualan) (model.Penjualan, error) {
		data.CreatedAt = time.Now()
		data.UpdatedAt 	= time.Now()
		err := data.CreatePenjualan(repository.Mysql.DB)

		return data, err
}

// 
func GetPenjualan() ([]model.Penjualan, error) {
		var penjualan model.Penjualan
		return penjualan.GetAll(repository.Mysql.DB)
}

// 
func GetPenjualanByID(id uint64) (model.Penjualan, error) {
		penjualan := model.Penjualan{
				ID: id,
		}
		return penjualan.GetPByID(repository.Mysql.DB)
}