package main

import (
	// "fmt"
	"log"
	// "projek/toko-retail/repository/config"

	"projek/toko-retail/services"

	"github.com/gofiber/fiber/v2"
	// "go.mongodb.org/mongo-driver/mongo"
)

// var client *mongo.Client


func main() {

	// client = repository.InitMongoDB()
	// _ = client

	// buat fiber app baru
	app := fiber.New()

	// initial route
	services.RouteSetup(app)

	// dbClient, err := cnfig.InitMongoDB(
	// if err!= nil {
    //     log.Fatal(err)
    // }

	// services.RouteSetup(app, dbClient)

	// fiber jalan di http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}

