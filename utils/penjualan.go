package utils

import (
	"fmt"
	"projek/toko-retail/model"
	repository "projek/toko-retail/repository/config"
	"time"
)

// function untuk buat kode invoice secara otomatis
func GenerateInvoice(id uint) string {
	invoice := fmt.Sprintf("INV - %d", id)
	return invoice
}

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