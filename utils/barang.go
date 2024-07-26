package utils

import (
	"fmt"
	"projek/toko-retail/model"
	repository "projek/toko-retail/repository/config"
	"strconv"
	"time"
)

// function untuk membuat data barang baru
func CreateBarang(data model.Barang) (model.CreateB, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	if data.CreatedBy == "" {
		data.CreatedBy = "SYSTEM"
	}

	data.Create(repository.Mysql.DB)
	if data.TipeBarang == "MAKANAN" {
			data.KodeBarang = fmt.Sprintf("MA-%v", strconv.FormatUint(uint64(data.ID), 10))
	} else if data.TipeBarang == "MINUMAN" {
			data.KodeBarang = fmt.Sprintf("MI-%v", strconv.FormatUint(uint64(data.ID), 10))
	} else {
			data.KodeBarang = fmt.Sprintf("L-%v", strconv.FormatUint(uint64(data.ID), 10))
	}

	data.Update(repository.Mysql.DB)

	histori, err := GetASK(data.ID)
	if err != nil {
			return model.CreateB{}, err
	}

	createB := model.CreateB{
			ID:				data.ID,
			KodeBarang: 	data.KodeBarang,
			Nama: 			data.Nama,
			HargaPokok: 	data.HargaPokok,
			HargaJual: 		data.HargaJual,
			TipeBarang: 	data.TipeBarang,
			Stok: 			data.Stok,
			Histori: 		histori,
	}

	return createB, nil
}

// function untuk mendapatkan list barang yang ada
func GetBarang() ([]model.Barang, error) {
	var barang model.Barang
	return barang.GetAll(repository.Mysql.DB)
}

// function untuk mendapatkan data barang berdasarkan ID barangnya
func GetBarangByID(id uint64) (model.Details, error) {
		barang := model.Barang{
				ID : id,
		}
		barang, err := barang.GetByID(repository.Mysql.DB)
		if err != nil {
			return model.Details{}, err
		}

		histori, err := GetASKMByIDBarang(barang.ID)
		if err != nil {
				return model.Details{}, err
		}

		details := model.Details{
				ID: barang.ID,
				KodeBarang: barang.KodeBarang,
				Nama: barang.Nama,
				HargaPokok: barang.HargaJual,
				HargaJual: barang.HargaJual,
				TipeBarang: barang.TipeBarang,
				Stok: barang.Stok,
				Model: barang.Model,
				Histori: histori,
		}
		return details, nil
}

// function untuk update data barang
func UpdateBarang(id uint, barang model.Barang) (model.Barang, error) {
		barang.ID = uint64(id)
		err := barang.Update(repository.Mysql.DB)
		return barang, err
}

func DeleteBarang(id uint64) error {
		barang := model.Barang{
				ID: id,
		}
		return barang.Delete(repository.Mysql.DB)
}

// {
// 	"nama_barang":"nasi",
// 	"harga_pokok":3000,
// 	"harga_jual":5000,
// 	"tipe_barang":"MAKANAN",
// 	"stok":20
//   }

// "data": {
//     "id": 1,
//     "kode_invoice": "INV/1",
//     "nama pembeli": "asep",
//     "subtotal": 20000,
//     "kode_diskon": null,
//     "diskon": 0,
//     "total": 20000,
//     "created_at": "2023-10-31 00:00:00",
//     "updated_at": "2023-10-31 00:00:00",
//     "deleted_at": null,
//     "created_by": "SYSTEM",
//     "item_penjualan": [
//       {
//         "kode_barang": "MI-1",
//         "jumlah": 1,
//         "subtotal": 10000,
//         "created_at": "2023-10-31 00:00:00",
//         "updated_at": "2023-10-31 00:00:00",
//         "deleted_at": null
//       },