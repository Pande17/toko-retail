package utils

import (
	"fmt"
	"projek/toko-retail/model"
	repository "projek/toko-retail/repository/config"
	"projek/toko-retail/repository/modelfunc"
	"strconv"
	"time"
)

// function untuk membuat data barang baru
func CreateBarang(data model.Barang) (model.CreateB, error) {
	// Initialize repository.Barang with model.Barang data
	repoBarang := modelfunc.Barang{
		Barang: data,
	}

	// Set timestamps and default CreatedBy if not provided
	repoBarang.CreatedAt = time.Now()
	repoBarang.UpdatedAt = time.Now()
	if repoBarang.CreatedBy == "" {
		repoBarang.CreatedBy = "SYSTEM"
	}

	// Create new barang record in the database
	err := repoBarang.Create(repository.Mysql.DB)
	if err != nil {
		return model.CreateB{}, err
	}

	// Set KodeBarang based on TipeBarang and newly created ID
	if repoBarang.TipeBarang == "MAKANAN" {
		repoBarang.KodeBarang = fmt.Sprintf("MA-%v", strconv.FormatUint(repoBarang.ID, 10))
	} else if repoBarang.TipeBarang == "MINUMAN" {
		repoBarang.KodeBarang = fmt.Sprintf("MI-%v", strconv.FormatUint(repoBarang.ID, 10))
	} else {
		repoBarang.KodeBarang = fmt.Sprintf("L-%v", strconv.FormatUint(repoBarang.ID, 10))
	}

	// Update barang record with the new KodeBarang
	err = repoBarang.Update(repository.Mysql.DB)
	if err != nil {
		return model.CreateB{}, err
	}

	// Fetch histori data for the newly created barang
	histori, err := GetASK(repoBarang.ID)
	if err != nil {
		return model.CreateB{}, err
	}

	// Prepare the CreateB response with updated data and histori
	createB := model.CreateB{
		ID:         repoBarang.ID,
		KodeBarang: repoBarang.KodeBarang,
		Nama:       repoBarang.Nama,
		HargaPokok: repoBarang.HargaPokok,
		HargaJual:  repoBarang.HargaJual,
		TipeBarang: repoBarang.TipeBarang,
		Stok:       repoBarang.Stok,
		CreatedBy:  repoBarang.CreatedBy,
		Histori:    histori,
	}

	return createB, nil
}

// function untuk mendapatkan list barang yang ada
func GetBarang() ([]model.Barang, error) {
	var barang modelfunc.Barang
	return barang.GetAll(repository.Mysql.DB)
}

// function untuk mendapatkan data barang berdasarkan ID barangnya
func GetBarangByID(id uint64) (model.Details, error) {
	barang := modelfunc.Barang{
		Barang: model.Barang{
			ID: id,
		},
	}
	barangModel, err := barang.GetByID(repository.Mysql.DB)
	if err != nil {
		return model.Details{}, err
	}

	histori, err := GetASKMByIDBarang(barangModel.ID)
	if err != nil {
		return model.Details{}, err
	}

	details := model.Details{
		ID:         barangModel.ID,
		KodeBarang: barangModel.KodeBarang,
		Nama:       barangModel.Nama,
		HargaPokok: barangModel.HargaPokok,
		HargaJual:  barangModel.HargaJual,
		TipeBarang: barangModel.TipeBarang,
		Stok:       barangModel.Stok,
		Model:      barangModel.Model,
		Histori:    histori,
	}
	return details, nil
}

// function untuk update data barang
func UpdateBarang(id uint, barang model.Barang) (model.Barang, error) {
	repoBarang := modelfunc.Barang{
		Barang: model.Barang{
			ID:         uint64(id),
			KodeBarang: barang.KodeBarang,
			Nama:       barang.Nama,
			HargaPokok: barang.HargaPokok,
			HargaJual:  barang.HargaJual,
			TipeBarang: barang.TipeBarang,
			Stok:       barang.Stok,
			Model:      barang.Model,
			CreatedBy:  barang.CreatedBy,
		},
	}
	err := repoBarang.Update(repository.Mysql.DB)
	return repoBarang.Barang, err
}

func DeleteBarang(id uint64) error {
	barang := modelfunc.Barang{
		Barang: model.Barang{
			ID: id,
		},
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