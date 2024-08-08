package controller

import (
	"fmt"
	"projek/toko-retail/model"
	"projek/toko-retail/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

	func CreateKodeDiskon(c *fiber.Ctx) error {
		type AddKodeDiskon struct {
			Kode_diskon string  `json:"kode_diskon"`
			Amount      float64 `json:"amount"`
			Type        string  `json:"type"`
		}

		req := new(AddKodeDiskon)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(map[string]any{
					"message": "Invalid Body",
				})

		}

		diskon, errDiskon := utils.CreateKodeDiskon(model.Diskon{
			KodeDiskon: req.Kode_diskon,
			Amount:     req.Amount,
			Type:       req.Type,
		})

		if errDiskon != nil {
			logrus.Printf("Terjadi error : %s\n", errDiskon.Error())
			return c.Status(fiber.StatusInternalServerError).
				JSON(map[string]any{
					"message": "Server Error",
				})
		}

		return c.Status(fiber.StatusOK).
			JSON(map[string]any{
				"data": diskon,
			})

	}

	func GetKodeDiskon(c *fiber.Ctx) error {
		dataDiskon, err := utils.GetDiskon()

		if err != nil {
			logrus.Error("Gagal dalam mengambil list Diskon: ", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(
				map[string]any{
					"message": "Server Error",
				},
			)
		}
		return c.Render("admin/kode-diskon", fiber.Map{
			"data": dataDiskon,    
			"title": "Daftar Diskon",
	})
	}

	func GetByCode(c *fiber.Ctx) error {
		DiskonCode := c.Query("kode-diskon")
		SubtotalStr := c.Query("subtotal")
		
		fmt.Println("Received kode-diskon:", DiskonCode)
		fmt.Println("Received subtotal:", SubtotalStr)
	
		if DiskonCode == "" {
			fmt.Println("kode-diskon is empty")
			return c.Status(fiber.StatusBadRequest).JSON(
				map[string]any{
					"message": "Invalid kode-diskon",
				},
			)
		}
	
		dataDiskon, err := utils.GetDiskonByCode(DiskonCode)
		if err != nil {
			fmt.Println("Error retrieving discount code:", err)
			if err.Error() == "record not found" {
				return c.Status(fiber.StatusNotFound).JSON(
					map[string]any{
						"message": "Discount Code not found",
					},
				)
			}
			return c.Status(fiber.StatusInternalServerError).JSON(
				map[string]any{
					"message": "Server Error",
				},
			)
		}
	
		var response fiber.Map
		if SubtotalStr != "" {
			subtotal, err := strconv.ParseFloat(SubtotalStr, 64)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(
					map[string]any{
						"message": "Invalid Subtotal",
					},
				)
			}
	
			var finalAmount float64
			if dataDiskon.Type == "PERCENT" {
				finalAmount = subtotal - (subtotal * (dataDiskon.Amount / 100))
			} else {
				finalAmount = subtotal - dataDiskon.Amount
			}
	
			response = fiber.Map{
				"subtotal": subtotal,
				"diskon":   dataDiskon.Amount,
				"total":    finalAmount,
			}
		} else {
			response = fiber.Map{
				"data": dataDiskon,
			}
		}
	
		return c.JSON(response)
	}
	
	
	

	func GetDiskonByID(c *fiber.Ctx) error {
		DiskonID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				map[string]any{
					"message": "Invalid ID",
				},
			)
		}

		dataDiskon, err := utils.GetDiskonByID(uint(DiskonID))
		if err != nil {
			if err.Error() == "record not found" {
				return c.Status(fiber.StatusNotFound).JSON(
					map[string]any{
						"message": "ID not found",
					},
				)
			}
		}

		return c.Status(fiber.StatusOK).JSON(
			map[string]any{
				"data": dataDiskon,
			},
		)
	}

	func UpdateCode(c *fiber.Ctx) error {
		DiskonID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid ID",
			})
		}

		var updatedDiskon model.Diskon
		if err := c.BodyParser(&updatedDiskon); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		dataDiskon, err := utils.UpdateDiskon(uint(DiskonID), updatedDiskon)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to update item",
			})
		}

		return c.Status(fiber.StatusOK).JSON(
			map[string]any{
				"data": dataDiskon,
			},
		)
	}

	func DeleteKode(c *fiber.Ctx) error {
		KodeID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				map[string]any{
					"message": "Invalid ID",
				},
			)
		}

		err = utils.DeleteKode(uint64(KodeID))
		if err != nil {
			if err.Error() == "record not found" {
				return c.Status(fiber.StatusNotFound).JSON(
					map[string]any{
						"message": "ID not found",
					},
				)
			}
		}

		return c.Status(fiber.StatusOK).JSON(map[string]any{
			"message": "Deleted Successfuly",
		},
		)
	}

	func ApplyDiskon(c *fiber.Ctx) error{
		type DiscountRequest struct {
			KodeDiskon 	string 	`json:"kode_diskon"`
			Subtotal	float64	`json:"subtotal"`
		}

		req := new(DiscountRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				map[string]any{
					"message":	"Invalid request body",
				},
			)
		}
		
		dataDiskon, err := utils.GetDiskonByCode(req.KodeDiskon)
		if err != nil {
			if err.Error() == "record not found" {
				return c.Status(fiber.StatusNotFound).JSON(
					map[string]any{
						"message":	"Discount code not found",
					},
				)
			}
			return c.Status(fiber.StatusInternalServerError).JSON(
				map[string]any{
					"message":	"Server error",
				},
			)
		}

		var finalAmount float64
		if dataDiskon.Type == "PERCENT" {
			finalAmount	= req.Subtotal - (req.Subtotal * (dataDiskon.Amount /100))
		} else {
			finalAmount = req.Subtotal - dataDiskon.Amount
		}

		return c.Status(fiber.StatusOK).JSON(
			map[string]any{
				"subtotal": 		req.Subtotal,
				"diskon":			dataDiskon.Amount,
				"total":			finalAmount,
			},
		)
	}