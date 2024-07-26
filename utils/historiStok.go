package utils

import (
	"projek/toko-retail/model"
	repository "projek/toko-retail/repository/config"
)

// function untuk membuat histori barang
func CreateHistoriBarang(p *model.Details, keterangan string, amount int, status string) (model.Histori, error) {
    histori := model.Histori{
        ID_barang: uint(p.ID),
        Amount: amount,
        Status: status,
        Keterangan: keterangan,
    }

    // Save the history record to the database
    err := histori.Create(repository.Mysql.DB)
    return histori, err
}

// function untuk membuat histori penjualan
func CreateHistoriPenjualan(p *model.CreateP, keterangan string, amount int, status string) (model.Histori, error) {
		Histori := model.Histori{
				ID_barang: uint(p.ID),
				Amount: amount,
				Status: status,
				Keterangan: keterangan,
		}

		return Histori, Histori.Create(repository.Mysql.DB)
}

// function untuk mendapatkan data dengan ID barang
func GetASKMByIDBarang(idb uint64) ([]model.HistoriASKM, error) {
		histori := model.Histori{
				ID_barang: uint(idb),
		}

		newHistory, err := histori.GetIDBarang(repository.Mysql.DB)
		if err != nil {
				return nil, err
		}

		var haskm []model.HistoriASKM
		for _,h := range newHistory {
				haskm = append(haskm, model.HistoriASKM{
						Amount: h.Amount,
						Status: h.Status,
						Keterangan: h.Keterangan,
						Model: h.Model,
				})
		}

		return haskm, err
}

// function untuk mendapatkan data
func GetASK(idb uint64) ([]model.HistoriASK, error) {
		histori := model.Histori{
				ID_barang: uint(idb),
		}

		newHistory, err := histori.GetIDBarang(repository.Mysql.DB)
		if err != nil {
				return nil, err
		}

		var hask []model.HistoriASK
		for _, h := range newHistory {
				hask = append(hask, model.HistoriASK{
					Amount: h.Amount,
					Status: h.Status,
					Keterangan: h.Keterangan,
				})
		}

		return hask, err
}