package controller

import (
	"projek/toko-retail/model"

	"github.com/gofiber/fiber/v2"
)

func GetBarang(ctx *fiber.Ctx) ([]model.Barang, error)  {
	// var barang model.Barang
	// return barang.GetAll(config.Mysql.DB)
	return
	
}

func GetBarangByID(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"Nigger": "Nig & Ger",
	})
}

func CreateBarang(c *fiber.Ctx) error {
	type AddBarangReq struct{
			Kode		string 	`json:"kode_barang"`
			Nama		string 	`json:"nama_barang"`
			HargaPokok	float64 `json:"harga_pokok"`
			HargaJual	float64 `json:"harga_jual"`
			Tipe		string 	`json:"tipe_barnag"`
			Histori		struct {
					Amount		uint	`json:"amount"`
					Status		string	`json:"status"`
					Keterangan	string	`json:"keterangan"`
			}`json:"histori_stok"`
	}
	req := new(AddBarangReq)
	if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).
					JSON(map[string]any{
						"message" : "Invalid Body",
				})
	}

	// barang, errCreateBarang := 
}

func IndexControllerRead(ctx *fiber.Ctx) error {
	
	return ctx.JSON(fiber.Map{
		"Title" : "Hello World",
	})
	
}