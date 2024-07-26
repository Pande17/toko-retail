package routes

import (
	"projek/toko-retail/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteSetup(r *fiber.App) {
	// r itu untuk 'route' 

	// Define routes
	retailGroup := r.Group("/toko")
	
	// Define main routes
	// retailGroup.Get("/",)

	// Define routes barang
	retailGroup.Get("/barang", controller.GetBarang)
	retailGroup.Get("/barang/:id", controller.GetBarangByID)
	retailGroup.Post("/barang", controller.CreateBarang)
	retailGroup.Put("/barang/:id", controller.UpdateBarang)
	retailGroup.Put("/barang/stok/:id", controller.UpdateStok)
	retailGroup.Delete("/barang/:id", controller.DeleteBarang)

	// Define routes penjualan
	retailGroup.Get("/penjualan", controller.GetPenjualan)
	retailGroup.Get("/penjualan/:id", controller.GetPenjualanByID)
	retailGroup.Post("/penjualan", controller.InsertPenjualanData)
	
	// Define routes kode diskon
	retailGroup.Get("/kode-diskon", controller.GetKodeDiskon)
	retailGroup.Get("/kode-diskon/:id", controller.GetDiskonByID)
	retailGroup.Get("/kode-diskon/get-by-code", controller.GetByCode)
	retailGroup.Post("/kode-diskon", controller.CreateKodeDiskon)
	retailGroup.Put("/kode-diskon/:id", controller.UpdateCode)
	retailGroup.Delete("kode-diskon/:id", controller.DeleteKode)
}