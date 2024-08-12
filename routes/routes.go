package routes

import (
	"projek/toko-retail/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteSetup(r *fiber.App) {
	// r itu untuk 'route'

	// Define routes
	retailGroup := r.Group("")

	// Define main routes
	// retailGroup.Get("/",)

	// Define routes barang
	retailGroup.Get("/barang", BarangHandler)
	retailGroup.Get("/keranjang", controller.KeranjangList) // ==== untuk keranjang, jangan lupa buat func
	retailGroup.Get("/barang/:id", controller.GetBarangByID)
	retailGroup.Post("/barang", controller.CreateBarang)
	retailGroup.Put("/barang/:id", controller.UpdateBarang)
	retailGroup.Put("/barang/stok/:id", controller.UpdateStok)
	retailGroup.Delete("/barang/:id", controller.DeleteBarang)

	// Define routes penjualan
	retailGroup.Get("/admin/penjualan", PenjualanHandler)
	retailGroup.Get("/admin/penjualan/:id", controller.GetPenjualanByID)
	retailGroup.Post("/penjualan", controller.InsertPenjualanData)

	// Define routes kode diskon
	retailGroup.Get("/admin/kode-diskon", controller.GetKodeDiskon)
	retailGroup.Get("/admin/kode-diskon/:id", controller.GetDiskonByID)
	retailGroup.Get("/kode-diskon/get-by-code", controller.GetByCode)
	retailGroup.Post("/kode-diskon", controller.CreateKodeDiskon)
	retailGroup.Put("/admin/kode-diskon/:id", controller.UpdateCode)
	retailGroup.Delete("/admin/kode-diskon/:id", controller.DeleteKode)

	retailGroup.Get("/test", func(c *fiber.Ctx) error {
		return c.Render("test", fiber.Map{
			"title": "Test Page",
		})
	})

	retailGroup.Get("/admin/dashboard", controller.AdminGetBarang)

	retailGroup.Get("/checkout", func(c *fiber.Ctx) error {
		return nil
	})

}
