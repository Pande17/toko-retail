package utils

import (
	"fmt"
	"projek/toko-retail/model"
	repository "projek/toko-retail/repository/config"
	"time"
)

// function untuk buat kode invoice secara otomatis
func GenerateInvoice(id uint64) string {
	invoice := fmt.Sprintf("INV/%d", id)
	return invoice
}

// function untuk memasukkan data penjualan ke DB
func InsertPenjualanData(data model.Penjualan) (model.Penjualan, error) {
		data.CreatedAt = time.Now()
		data.UpdatedAt 	= time.Now()

		// mendapatkan diskon berdasarkan kode diskon
		if data.Kode_diskon != "" {
			diskon, err := GetDiskonByCode(data.Kode_diskon)
			if err != nil {
				return data, err
			}

			// hitung diskon
			var diskonAmount float64
			if diskon.Type == "PERCENT" {
				diskonAmount = data.Subtotal * (diskon.Amount / 100)
			}else {
				diskonAmount = diskon.Amount
			}

			// terapkan diskon pada subtotal
			data.Diskon = diskonAmount
			data.Total = data.Subtotal - data.Diskon
		} else {
			data.Diskon = 0 
			data.Total = data.Subtotal
		}

		// masukkan generate invoice
		data.CreatePenjualan(repository.Mysql.DB)
		data.Kode_invoice = GenerateInvoice(data.ID) 
		data.Update(repository.Mysql.DB)

		// masukkan diskon dan pengurangan harga setelah diskon
		// data.CreateDiskon(repository.Mysql.DB)
		// data.Total = ApplyDiskon(data.ID)
		

		return data, nil
}

//  function untuk mendapatkan data penjualan
func GetPenjualan() ([]model.Penjualan, error) {
		var penjualan model.Penjualan
		return penjualan.GetAll(repository.Mysql.DB)
}

// function untuk mendapatkan data penjualan berdasarkan ID
func GetPenjualanByID(id uint64) (model.Penjualan, error) {
		penjualan := model.Penjualan{
				ID: id,
		}
		return penjualan.GetPByID(repository.Mysql.DB)
}