package controller

import "github.com/gofiber/fiber/v2"

func UserControllerRead(ctx *fiber.Ctx) error {
	
	return ctx.JSON(fiber.Map{
		"kode_barang" : "MA-2192jsd",
	})
	
}