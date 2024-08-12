package controller

import (
	"projek/toko-retail/model"
	"projek/toko-retail/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func CreateBarang(c *fiber.Ctx) error {
	type AddBarangReq struct {
		Kode       string  `json:"kode_barang"`
		Nama       string  `json:"nama_barang"`
		HargaPokok float64 `json:"harga_pokok"`
		HargaJual  float64 `json:"harga_jual"`
		Tipe       string  `json:"tipe_barang"`
		Stok       uint    `json:"stok"`
		CreatedBy  string  `json:"created_by"`
		Histori    struct {
			Amount     uint   `json:"amount"`
			Status     string `json:"status"`
			Keterangan string `json:"keterangan"`
		} `json:"histori_stok"`
	}
	req := new(AddBarangReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{
				"message": "Invalid Body",
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

func GetBarang(c *fiber.Ctx) ([]model.Barang, error) {
	dataBarang, err := utils.GetBarang()
	if err != nil {
		logrus.Error("Gagal dalam mengambil list Barang: ", err.Error())
		return nil, err
	}
	return dataBarang, nil
}

func GetJSONBarang(c *fiber.Ctx) error {
	dataBarang, err := utils.GetBarang()
	if err != nil {
		logrus.Error("Gagal dalam mengambil list Barang: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "server error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"barang": dataBarang,
		},
	)
}

func AdminGetBarang(c *fiber.Ctx) error {
	dataBarang, err := utils.GetBarang()
	if err != nil {
		logrus.Error("Gagal dalam mengambil list Barang: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "server error",
			},
		)
	}
	return c.Render("admin/dashboard", fiber.Map{
		"data":  dataBarang,
		"title": "Daftar Barang",
	})
}

func GetBarangByID(c *fiber.Ctx) error {
	barangID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "Invalid ID",
			},
		)
	}

	dataBarang, err := utils.GetBarangByID(uint64(barangID))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",
				},
			)
		}
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data": dataBarang,
		},
	)
}

func UpdateBarang(c *fiber.Ctx) error {
	barangID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	var updatedBarang model.Barang
	if err := c.BodyParser(&updatedBarang); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	dataBarang, err := utils.UpdateBarang(uint(barangID), updatedBarang)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update item",
		})
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"id": dataBarang.ID,
		},
	)
}

func DeleteBarang(c *fiber.Ctx) error {
	barangID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "Invalid ID",
			},
		)
	}

	err = utils.DeleteBarang(uint64(barangID))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",
				},
			)
		}
	}

	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"message": "Deleted Successfully",
	},
	)
}

func UpdateStok(c *fiber.Ctx) error {
	barangID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	var requestData struct {
		Stok        uint          `json:"stok"`
		HistoriStok model.Histori `json:"histori_stok"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Retrieve the existing Barang to update it
	existingBarang, err := utils.GetBarangByID(uint64(barangID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve item",
		})
	}

	// Update only the fields provided in the request
	existingBarang.Stok = requestData.Stok

	// Update the Barang in the database
	updatedBarang := model.Barang{
		ID:         existingBarang.ID,
		KodeBarang: existingBarang.KodeBarang,
		Nama:       existingBarang.Nama,
		HargaPokok: existingBarang.HargaPokok,
		HargaJual:  existingBarang.HargaJual,
		TipeBarang: existingBarang.TipeBarang,
		Stok:       existingBarang.Stok,
	}
	updatedBarang, err = utils.UpdateBarang(uint(existingBarang.ID), updatedBarang)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update item",
		})
	}

	// Create the history record
	newHistori, err := utils.CreateHistoriBarang(&existingBarang, requestData.HistoriStok.Keterangan, requestData.HistoriStok.Amount, requestData.HistoriStok.Status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create history record",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":          updatedBarang.ID,
		"kode_barang": updatedBarang.KodeBarang,
		"stok":        updatedBarang.Stok,
		"histori_stok": map[string]interface{}{
			"amount":     newHistori.Amount,
			"status":     newHistori.Status,
			"keterangan": newHistori.Keterangan,
		},
	})
}

func KeranjangList(c *fiber.Ctx) error {
	return c.Render("barang/keranjang", fiber.Map{
		"title": "Keranjang Belanja",
	})
}
