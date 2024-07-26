package main

import (
	// "fmt"
	"log"
	// "projek/toko-retail/repository/config"

	repository "projek/toko-retail/repository/config"
	"projek/toko-retail/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// initial database
	repository.InitDB()

	// buat fiber app baru
	app := fiber.New()

	// initial route
	routes.RouteSetup(app)

	// fiber jalan di http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}

