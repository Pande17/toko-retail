package controller

import (
	"projek/toko-retail/model"
	"projek/toko-retail/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func GetBarang(ctx *fiber.Ctx) ([]model.Barang, error)  {
	// var barang model.Barang
	// return barang.GetAll(config.Mysql.DB)
	return nil, nil
	
}

func GetBarangByID(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"Hello": "Test",
	})
}

func CreateBarang(c *fiber.Ctx) error {
	type AddBarangReq struct{
			Kode		string 	`json:"kode_barang"`
			Nama		string 	`json:"nama_barang"`
			HargaPokok	float64 `json:"harga_pokok"`
			HargaJual	float64 `json:"harga_jual"`
			Tipe		string 	`json:"tipe_barang"`
			Stok       uint         `json:"stok"`
			CreatedBy  string       `json:"created_by"`
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

	barang, errCreateBarang := utils.CreateBarang(model.Barang{
		KodeBarang: req.Kode,
		Nama:       req.Nama,
		HargaPokok: req.HargaPokok,
		HargaJual:  req.HargaJual,
		TipeBarang: req.Tipe,
		Stok:       req.Stok,
		CreatedBy:  req.CreatedBy,
	})

	utils.CreateHistoriBarang(&model.Details{
		ID:         barang.ID,
		KodeBarang: req.Kode,
		Nama:       req.Nama,
		HargaPokok: req.HargaPokok,
		HargaJual:  req.HargaJual,
		TipeBarang: req.Tipe,
		Stok:       req.Stok,
		CreatedBy:  req.CreatedBy,
		Histori:    []model.HistoriASKM{},
	}, req.Histori.Keterangan, int(req.Stok), req.Histori.Status)

	if errCreateBarang != nil {
		logrus.Printf("Terjadi error : %s\n", errCreateBarang.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(map[string]any{
			"id":          barang.ID,
			"kode_barang": barang.KodeBarang,
		})
}

func IndexControllerRead(ctx *fiber.Ctx) error {
	
	return ctx.JSON(fiber.Map{
		"Title" : "Hello World",
	})
	
}