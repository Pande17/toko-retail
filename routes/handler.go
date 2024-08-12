package routes

import (
	"projek/toko-retail/controller"

	"github.com/gofiber/fiber/v2"
)

func BarangHandler(c *fiber.Ctx) error {
	dataBarang, err := controller.GetBarang(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get data")
	}

	return c.Render("barang/list", fiber.Map{
		"data":  dataBarang,
		"title": "Daftar Barang",
	})
}

func PenjualanHandler(c *fiber.Ctx) error {
	dataPenjualan := controller.GetPenjualan(&fiber.Ctx{})

	return c.Render("admin/penjualan", fiber.Map{
		"data":  dataPenjualan,
		"title": "Daftar Penjualan",
	})
}
