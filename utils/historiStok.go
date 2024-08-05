package utils

import (
	"projek/toko-retail/model"
	repository "projek/toko-retail/repository/config"
	"projek/toko-retail/repository/modelfunc"
)

// function untuk membuat histori barang
func CreateHistoriBarang(p *model.Details, keterangan string, amount int, status string) (model.Histori, error) {
	histori := modelfunc.Histori{
		Histori: model.Histori{
			ID_barang:   uint(p.ID),
			Amount:      amount,
			Status:      status,
			Keterangan:  keterangan,
		},
	}

	// Save the history record to the database
	err := histori.Create(repository.Mysql.DB)
	if err != nil {
		return model.Histori{}, err
	}

	return histori.Histori, nil
}

// function untuk membuat histori penjualan
func CreateHistoriPenjualan(p *model.CreateP, keterangan string, amount int, status string) (model.Histori, error) {
	histori := modelfunc.Histori{
		Histori: model.Histori{
			ID_barang:   uint(p.ID),
			Amount:      amount,
			Status:      status,
			Keterangan:  keterangan,
		},
	}

	err := histori.Create(repository.Mysql.DB)
	if err != nil {
		return model.Histori{}, err
	}

	return histori.Histori, nil
}

// function untuk mendapatkan data dengan ID barang
func GetASKMByIDBarang(idb uint64) ([]model.HistoriASKM, error) {
	histori := modelfunc.Histori{
		Histori: model.Histori{
			ID_barang: uint(idb),
		},
	}

	newHistory, err := histori.GetIDBarang(repository.Mysql.DB)
	if err != nil {
		return nil, err
	}

	var haskm []model.HistoriASKM
	for _, h := range newHistory {
		haskm = append(haskm, model.HistoriASKM{
			Amount:     h.Amount,
			Status:     h.Status,
			Keterangan: h.Keterangan,
			Model:      h.Model,
		})
	}

	return haskm, nil
}

// function untuk mendapatkan data
func GetASK(idb uint64) ([]model.HistoriASK, error) {
	histori := modelfunc.Histori{
		Histori: model.Histori{
			ID_barang: uint(idb),
		},
	}

	newHistory, err := histori.GetIDBarang(repository.Mysql.DB)
	if err != nil {
		return nil, err
	}

	var hask []model.HistoriASK
	for _, h := range newHistory {
		hask = append(hask, model.HistoriASK{
			Amount:     h.Amount,
			Status:     h.Status,
			Keterangan: h.Keterangan,
		})
	}

	return hask, nil
}