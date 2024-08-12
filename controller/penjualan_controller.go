package controller

import (
	"projek/toko-retail/model"
	"projek/toko-retail/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	logrus "github.com/sirupsen/logrus"
)

func InsertPenjualanData(c *fiber.Ctx) error {
	type AddPenjualanReq struct {
		ID            uint64  `json:"id"`
		KodeInvoice   string  `json:"kode_invoice"`
		NamaPembeli   string  `json:"nama_pembeli"`
		Subtotal      float64 `json:"subtotal"`
		KodeDiskon    string  `json:"kode_diskon"`
		Diskon        float64 `json:"diskon"`
		Total         float64 `json:"total"`
		CreatedBy     string  `json:"created_by"`
		ItemPenjualan []struct {
			Kode   string `json:"kode_barang"`
			Jumlah uint   `json:"jumlah"`
		} `json:"item_penjualan"`
	}
	req := new(AddPenjualanReq)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]interface{}{
				"message": "invalid Body",
			})
	}

	penjualan := model.Penjualan{
		ID:           req.ID,
		Kode_invoice: req.KodeInvoice,
		Nama_pembeli: req.NamaPembeli,
		Subtotal:     req.Subtotal,
		Kode_diskon:  req.KodeDiskon,
		Diskon:       req.Diskon,
		Total:        req.Total,
		Created_by:   req.CreatedBy,
	}

	_, errInsertPenjualan := utils.InsertPenjualanData(penjualan)
	if errInsertPenjualan != nil {
		logrus.Printf("Terjadi error : %s\n", errInsertPenjualan.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]interface{}{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(map[string]interface{}{
			"message": "Berhasil Menambahkan Barang",
		})
}

func GetPenjualan(c *fiber.Ctx) error {
	dataPenjualan, err := utils.GetPenjualan()
	if err != nil {
		logrus.Error("Gagal dalam mengambil list penjualan :", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]interface{}{
				"message": "Server Error",
			},
		)
	}

	if dataPenjualan != nil {
		logrus.Info("Data Penjualan yang diterima: ", dataPenjualan)
		logrus.Info("Jumlah item dalam data penjualan: ", len(dataPenjualan))
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"penjualan": dataPenjualan,
		},
	)
}

func GetPenjualanByID(c *fiber.Ctx) error {
	penjualanID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]interface{}{
				"message": "Invalid ID",
			},
		)
	}

	dataPenjualan, err := utils.GetPenjualanByID(uint64(penjualanID))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]interface{}{
					"message": "ID not found",
				},
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]interface{}{
				"message": "Server Error",
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]interface{}{
			"data":    dataPenjualan,
			"message": "success",
		},
	)
}
