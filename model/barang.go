package model

type Barang struct {
	ID         uint64  `gorm:"primarykey" json:"id"`
	KodeBarang string  `json:"kode_barang"`
	Nama       string  `json:"nama_barang"`
	HargaPokok float64 `json:"harga_pokok"`
	HargaJual  float64 `json:"harga_jual"`
	TipeBarang string  `json:"tipe_barang"`
	Stok       uint    `json:"stok"`
	Model
	CreatedBy string `json:"created_by"`
}

type Details struct {
	ID         uint64  `gorm:"primarykey" json:"id"`
	KodeBarang string  `json:"kode_barang"`
	Nama       string  `json:"nama_barang"`
	HargaPokok float64 `json:"harga_pokok"`
	HargaJual  float64 `json:"harga_jual"`
	TipeBarang string  `json:"tipe_barang"`
	Stok       uint    `json:"stok"`
	Model
	CreatedBy string        `json:"created_by"`
	Histori   []HistoriASKM `gorm:"foreignKey:ID_Barang" json:"histori_stok"`
}

type CreateB struct {
	ID         uint64       `gorm:"primarykey" json:"id"`
	KodeBarang string       `json:"kode_barang"`
	Nama       string       `json:"nama_barang"`
	HargaPokok float64      `json:"harga_pokok"`
	HargaJual  float64      `json:"harga_jual"`
	TipeBarang string       `json:"tipe_barang"`
	Stok       uint         `json:"stok"`
	CreatedBy  string       `json:"created_by"`
	Histori    []HistoriASK `gorm:"foreignKey:ID_Barang" json:"histori_stok"`
}

