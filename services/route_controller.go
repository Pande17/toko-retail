package services

import (
	// "context"
	// "log"

	// "projek/toko-retail/model"

	"projek/toko-retail/controller"

	"github.com/gofiber/fiber/v2"
	// "go.mongodb.org/mongo-driver/mongo"
)

// var client *mongo.Client

func RouteSetup(r *fiber.App) {
	
	// Define routes
	// r.Get("/", )
	r.Get("/user", controller.UserControllerRead)
	
	// app.Get("/barang", func(c *fiber.Ctx) error {
	// 	collection := client.Database("tokoRetail").Collection("barang")
		
	// 	// query untuk melihat semua barang
	// 	cur, err := collection.Find(context.TODO(), map[string]interface{}{})
	// 	if err!= nil {
	// 		return c.SendStatus(fiber.StatusInternalServerError)
    //     }
	// 	defer cur.Close(context.TODO())
	// 	var Barang map[string]interface{}
	// 	data := []model.Barang{}
	// 	for cur.Next(context.TODO()) {
	// 		var barang map[string]interface{}
	// 		if err := cur.Decode(&barang); err != nil {
	// 			return c.SendStatus(fiber.StatusInternalServerError)
	// 		}
	// 		Barang = append(data, barang)
	// 	}

	// 	return c.Json(Barang)
	// })
	
}