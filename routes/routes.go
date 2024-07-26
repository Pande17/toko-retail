package routes

import (
	"projek/toko-retail/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteSetup(r *fiber.App) {
	// r itu untuk 'route' 

	// Define routes
	retailGroup := r.Group("/toko")
	
	
	retailGroup.Get("/",)
	// Define routes barang
	retailGroup.Get("/barang",)
	retailGroup.Get("/barang/:id", controller.GetBarangByID)
	retailGroup.Post("/barang", )
	retailGroup.Put("/barang/:id", )
	retailGroup.Put("/barang/stok/:id", )
	retailGroup.Delete("/barang/:id", )

	// Define routes penjualan
	retailGroup.Get("/penjualan", )
	retailGroup.Get("/penjualan/:id", )
	retailGroup.Post("/penjualan", )
	
	// Define routes kode diskon
	retailGroup.Get("/kode-diskon", )
	retailGroup.Get("/kode-diskon/:id", )
	retailGroup.Get("/kode-diskon/get-by-code", )
	retailGroup.Post("/kode-diskon", )
	retailGroup.Put("/kode-diskon/:id", )
	retailGroup.Delete("kode-diskon/:id", )
}