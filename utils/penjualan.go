package utils

import (
	"fmt"
	"projek/toko-retail/model"
	repository "projek/toko-retail/repository/config"
	"projek/toko-retail/repository/modelfunc"
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
	data.UpdatedAt = time.Now()

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
		} else {
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
	data.Kode_invoice = GenerateInvoice(data.ID)

	// convert model.Penjualan to modelfunc.Penjualan
	penjualan := modelfunc.Penjualan{
		Penjualan: data,
	}

	err := penjualan.CreatePenjualan(repository.Mysql.DB)
	if err != nil {
		return data, err
	}

	err = penjualan.Update(repository.Mysql.DB)
	if err != nil {
		return data, err
	}

	return penjualan.Penjualan, nil
}

//  function untuk mendapatkan data penjualan
func GetPenjualan() ([]model.Penjualan, error) {
		var penjualan modelfunc.Penjualan
		penjualanList, err := penjualan.GetAll(repository.Mysql.DB)
	if err != nil {
		return nil, err
	}

	// Convert []modelfunc.Penjualan to []model.Penjualan
	result := make([]model.Penjualan, len(penjualanList))
	for i, pj := range penjualanList {
		result[i] = pj.Penjualan
	}

	return result, nil
}

// function untuk mendapatkan data penjualan berdasarkan ID
func GetPenjualanByID(id uint64) (model.Penjualan, error) {
	penjualan := modelfunc.Penjualan{
		Penjualan: model.Penjualan{
			ID: id,
		},
	}
	result, err := penjualan.GetPByID(repository.Mysql.DB)
	if err != nil {
		return model.Penjualan{}, err
	}
	return result.Penjualan, nil
}