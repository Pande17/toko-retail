package model

type Histori struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	ID_barang  uint   `json:"id_barang"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Keterangan string `json:"keterangan"`
	Model
}

type HistoriASK struct {
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Keterangan string `json:"keterangan"`
}

type HistoriASKM struct {
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Keterangan string `json:"keterangan"`
	Model
}
