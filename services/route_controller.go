package services

import (
	// "context"
	// "log"

	// "projek/toko-retail/model"

	"projek/toko-retail/controller"
	"github.com/gofiber/fiber/v2"
)

func RouteSetup(r *fiber.App) {
	
	// Define routes
	// r.Get("/", )
	r.Get("/user", controller.UserControllerRead)
	
}