package controller

import (
	"projek/toko-retail/model"
	"projek/toko-retail/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// Niggers
func InsertPenjualanData(c *fiber.Ctx) error {
		type AddPenjualanReq struct {
				NamaPembeli 	string	`json:"nama"`
				ID				uint64	`json:"id"`
				ItemPenjualan	[]struct {
						Kode	string	`json:"kode_barang"`
						Jumlah	uint	`json:"jumlah"`
				} `json:"item_penjualan"`
		}
		req := new(AddPenjualanReq)

		if err := c.BodyParser(req); err != nil {
				return c.Status(fiber.StatusBadRequest).
						JSON(map[string]any{
								"message": "invalid Body",
						})
		}

		_, errInsertpPenjualan := utils.InsertPenjualanData(model.Penjualan{
				Nama_pembeli: 	req.NamaPembeli,
				ID: 			req.ID,
		})	

		if errInsertpPenjualan != nil {
				logrus.Printf("Terjadi error : %s\n", errInsertpPenjualan.Error())
				return c.Status(fiber.StatusInternalServerError).
						JSON(map[string]any{
								"message": "Server Error",
						})
		}

		return c.Status(fiber.StatusOK).
				JSON(map[string]any{
						"message" : "Berhasil Menambahkan Barang",
				})
}

func GetPenjualan(c *fiber.Ctx)	error {
		dataPenjualan, err := utils.GetPenjualan()
		if err != nil {
				logrus.Error("Gagal dalam mengambil list penjualan :", err.Error())
				return c.Status(fiber.StatusInternalServerError).JSON(
						map[string]any{
								"message": "Server Error",
						},
				)
		}
		return c.Render("admin/penjualan", fiber.Map{
			"data": dataPenjualan,    
			"title": "Daftar Penjualan",
	})
}

func GetPenjualanByID(c *fiber.Ctx) error{
		penjualanID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(
					map[string]any{
							"message": "Invalid ID",
					},
				)
		}

		dataPenjualan, err := utils.GetPenjualanByID(uint64(penjualanID))
		if err != nil {
				if err.Error() == "record not found" {
						return c.Status(fiber.StatusNotFound).JSON(
								map[string]any{
										"message":	"ID not found",
								},
						)
				}
		}

		return c.Status(fiber.StatusOK).JSON(
				map[string]any{
						"data":		dataPenjualan,
						"message":	"success",
				},
		)
}