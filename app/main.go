package main

import (
	// "fmt"
	// "log"
	// "projek/toko-retail/repository/config"

	repository "projek/toko-retail/repository/config"
	"projek/toko-retail/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Warn("Cannot load env file, using system env")
	}
}

func main() {
	// Init environment
	InitEnv()

	// initial database
	repository.OpenDB()

	engine := html.New("./template", ".html") // Path to templates and file extension
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./template/home")
	app.Static("/src", "./src")

	// initial route
	routes.RouteSetup(app)

	// open fiber on http://localhost:3000
	err := app.Listen(":3000")
	if err != nil {
		logrus.Fatal(
			"Error on running fiber, ",
			err.Error())
	}
}
