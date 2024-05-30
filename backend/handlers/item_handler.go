package handlers

import (
	"fmt"
	"stock-manager/database"
	"stock-manager/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetItems(c *fiber.Ctx) error {
	db := database.Db
	var items []models.Item

	err := db.Find(&items).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": false,
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   items,
	})
}

func GetItem(c *fiber.Ctx) error {
	db := database.Db
	var item models.Item

	kode := c.Params("kode")

	err := db.First(&item, "kode = ?", kode).Error

	if err != nil && err == gorm.ErrRecordNotFound {
		return c.Status(404).JSON(fiber.Map{
			"status": false,
			"error":  "Item not found",
			"data":   nil,
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   item,
	})
}

func CreateItem(c *fiber.Ctx) error {
	db := database.Db
	item := new(models.Item)

	err := c.BodyParser(item)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": false,
		})
	}

	err = db.Create(&item).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": false,
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
	})
}

func UpdateItem(c *fiber.Ctx) error {
	type updateItem struct {
		Nama      string `json:"nama"`
		Jumlah    int    `json:"jumlah"`
		Deskripsi string `json:"deskripsi"`
		Status    bool   `json:"status"`
	}

	db := database.Db
	var item models.Item

	kode := c.Params("kode")

	err := db.Find(&item, "kode = ?", kode).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"status": false,
				"error":  "Item not found",
				"data":   nil,
			})
		}
	}

	var updateItemData updateItem
	err = c.BodyParser(&updateItemData)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": false,
		})
	}

	item.Nama = updateItemData.Nama
	item.Jumlah = uint(updateItemData.Jumlah)
	item.Deskripsi = updateItemData.Deskripsi
	item.Status = updateItemData.Status

	db.Save(&item)

	return c.JSON(fiber.Map{
		"status": true,
		"data":   item,
	})
}

func DeleteItem(c *fiber.Ctx) error {
	db := database.Db
	var item models.Item

	kode := c.Params("kode")

	err := db.Find(&item, "kode = ?", kode).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"status": false,
				"error":  "Item not found",
				"data":   nil,
			})
		}
	}

	item.Status = false

	db.Save(&item)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": fmt.Sprintf("Success delete item %s", kode),
	})
}
