package payload

import (
	"time"

	"gorm.io/gorm"
)

type ItemPenjualanRequest struct {
	KodeBarang string `json:"kode_barang"`
	Jumlah      uint   `json:"jumlah"`
	Subtotal    uint   `json:"subtotal"`
}

type AddPenjualanRequest struct {
	Nama_pembeli   string                 `json:"nama_pembeli" valid:"required, type(string)"`
	Subtotal       float64                `json:"subtotal" valid:"required"`
	Kode_diskon    string                 `json:"kode_diskon"`
	Total          float64                `json:"total" valid:"optional , type(float64)"`
	CreatedBy      string                 `json:"created_by" valid:"required, type(string)"`
	Item_Penjualan []ItemPenjualanRequest `json:"item_penjualan"`
}

type GetPenjualanRespone struct {
	ID           uint           `json:"id"`
	Kode_Invoice string         `json:"kode_invoice"`
	Nama_pembeli string         `json:"nama_pembeli"`
	Subtotal     float64        `json:"subtotal"`
	Kode_diskon  string         `json:"kode_diskon"`
	Diskon       float64        `json:"diskon"`
	Total        float64        `json:"total"`
	Created_at   time.Time      `json:"created_at"`
	Updated_at   time.Time      `json:"updated_at"`
	Deleted_at   gorm.DeletedAt `json:"deleted_at"`
	Created_By   string         `json:"created_by"`
}

type GetPenjualanDetailRespone struct {
	ID             uint                   `json:"id"`
	Kode_Invoice   string                 `json:"kode_invoice"`
	Nama_pembeli   string                 `json:"nama_pembeli"`
	Subtotal       float64                `json:"subtotal"`
	Kode_diskon    string                 `json:"kode_diskon"`
	Diskon         float64                `json:"diskon"`
	Total          float64                `json:"total"`
	Created_at     time.Time              `json:"created_at"`
	Updated_at     time.Time              `json:"updated_at"`
	Deleted_at     gorm.DeletedAt         `json:"deleted_at"`
	Created_By     string                 `json:"created_by"`
	Item_Penjualan []ItemPenjualanRespone `json:"item_penjualan"`
}
type ItemPenjualanRespone struct {
	KodeBarang string         `json:"kode_barang"`
	Jumlah      uint           `json:"jumlah"`
	Subtotal    float64        `json:"subtotal"`
	Created_At  time.Time      `json:"created_at"`
	Updated_At  time.Time      `json:"updated_at"`
	Deleted_at  gorm.DeletedAt `json:"deleted_at"`
}
