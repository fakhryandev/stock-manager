package handlers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"stock-manager/database"
	"stock-manager/models"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
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

	kode := c.Params("code")
	kode = strings.ToUpper(kode)

	err := db.First(&item, "kode = ?", &kode).Error

	if err != nil && err == gorm.ErrRecordNotFound {
		return c.Status(404).JSON(fiber.Map{
			"status": false,
			"error":  "Item tidak ditemukan",
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

	validate := validator.New()
	err = validate.Struct(item)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	item.Kode = strings.ToUpper(item.Kode)
	err = db.Create(&item).Error

	if err != nil {
		var psqlErr *pgconn.PgError
		if errors.As(err, &psqlErr) && psqlErr.Code == "23505" {
			return c.Status(400).JSON(fiber.Map{
				"status":  false,
				"message": "Gagal, kode item sudah ada.",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Sukses, data berhasil ditambahkan",
		"data":    item,
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

	kode := c.Params("code")
	kode = strings.ToUpper(kode)

	err := db.First(&item, "kode = ?", kode).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"status": false,
				"error":  "Item tidak ditemukan",
				"data":   nil,
			})
		}
	}

	var updateItemData updateItem
	err = c.BodyParser(&updateItemData)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
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

	kode := c.Params("code")
	kode = strings.ToUpper(kode)

	err := db.First(&item, "kode = ?", kode).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"status": false,
				"error":  "Item not found",
				"data":   nil,
			})
		}
	}

	db.Delete(&item)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": fmt.Sprintf("Success delete item %s", kode),
	})
}

func IncreaseItem(c *fiber.Ctx) error {
	type updateAmountItem struct {
		Jumlah uint `json:"jumlah"`
	}

	db := database.Db
	var item models.Item

	kode := c.Params("code")
	kode = strings.ToUpper(kode)

	err := db.First(&item, "kode = ?", kode).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"status": false,
				"error":  "Item tidak ditemukan",
				"data":   nil,
			})
		}
	}

	var updateAmountItemData updateAmountItem
	err = c.BodyParser(&updateAmountItemData)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	item.Jumlah = item.Jumlah + uint(updateAmountItemData.Jumlah)

	db.Save(&item)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Jumlah stok berhasil ditambahkan.",
		"data":    item,
	})
}

func DecreaseItem(c *fiber.Ctx) error {
	type updateAmountItem struct {
		Jumlah uint `json:"jumlah"`
	}

	db := database.Db
	var item models.Item

	kode := c.Params("code")
	kode = strings.ToUpper(kode)

	err := db.First(&item, "kode = ?", kode).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"status": false,
				"error":  "Item tidak ditemukan",
				"data":   nil,
			})
		}
	}

	var updateAmountItemData updateAmountItem
	err = c.BodyParser(&updateAmountItemData)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": false,
		})
	}

	amountNow := item.Jumlah - uint(updateAmountItemData.Jumlah)

	if int(amountNow) < 0 {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": "Total stok yang dikurangi tidak boleh kurang dari 0",
		})
	}

	item.Jumlah = amountNow

	db.Save(&item)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Jumlah stok berhasil dikurangi.",
		"data":    item,
	})
}

func ImportFile(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal mendapatkan file",
		})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal membuka file csv",
		})
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := readCSV(csvReader)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal membaca file csv",
		})
	}

	tx := database.Db.Begin()
	for _, record := range records[1:] {
		operation := "insert"
		jumlah, _ := strconv.Atoi(record[3])
		if record[4] == "edit" {
			operation = "edit"
		}

		if record[4] == "tambah" {
			operation = "tambah"
		}

		if record[4] == "kurang" {
			operation = "kurang"
		}

		if operation == "edit" {
			var item models.Item
			kode := record[0]
			kode = strings.ToUpper(kode)
			err := tx.First(&item, "kode = ?", kode).Error

			if err != nil {
				tx.Rollback()
				if err == gorm.ErrRecordNotFound {
					return c.Status(404).JSON(fiber.Map{
						"status": false,
						"error":  fmt.Sprintf("Item dengan kode %s tidak ditemukan. seluruh data dikembalikan", kode),
						"data":   nil,
					})
				}
			}

			item.Nama = record[1]
			item.Deskripsi = record[2]
			item.Jumlah = uint(jumlah)

			tx.Save(&item)

		} else if operation == "tambah" {
			var item models.Item
			kode := record[0]
			kode = strings.ToUpper(kode)

			err := tx.First(&item, "kode = ?", kode).Error

			if err != nil {
				tx.Rollback()
				if err == gorm.ErrRecordNotFound {
					return c.Status(404).JSON(fiber.Map{
						"status": false,
						"error":  fmt.Sprintf("Item dengan kode %s tidak ditemukan. seluruh data dikembalikan", kode),
						"data":   nil,
					})
				}
			}

			item.Jumlah = item.Jumlah + uint(jumlah)
			tx.Save(&item)

		} else if operation == "kurang" {
			var item models.Item
			kode := record[0]
			kode = strings.ToUpper(kode)

			err := tx.First(&item, "kode = ?", kode).Error

			if err != nil {
				tx.Rollback()
				if err == gorm.ErrRecordNotFound {
					return c.Status(404).JSON(fiber.Map{
						"status": false,
						"error":  fmt.Sprintf("Item dengan kode %s tidak ditemukan. seluruh data dikembalikan", kode),
						"data":   nil,
					})
				}
			}

			item.Jumlah = item.Jumlah - uint(jumlah)

			if int(item.Jumlah) < 0 {
				tx.Rollback()
				return c.Status(400).JSON(fiber.Map{
					"status": false,
					"error":  fmt.Sprintf("Total item dengan kode %s tidak boleh kurang dari 0. seluruh data dikembalikan", kode),
					"data":   nil,
				})
			}

			tx.Save(&item)
		} else {
			newItem := models.Item{
				Kode:      record[0],
				Nama:      record[1],
				Deskripsi: record[2],
				Jumlah:    uint(jumlah),
			}

			if err := tx.Create(&newItem).Error; err != nil {
				tx.Rollback()
				var psqlErr *pgconn.PgError
				if errors.As(err, &psqlErr) && psqlErr.Code == "23505" {
					return c.Status(400).JSON(fiber.Map{
						"status":  false,
						"message": "Gagal, kode item terduplikasi.",
					})
				}
				return c.Status(500).JSON(fiber.Map{
					"status":  false,
					"message": err.Error(),
				})
			}
		}

	}

	tx.Commit()

	return c.JSON(fiber.Map{
		"status": true,
	})
}

func readCSV(reader *csv.Reader) ([][]string, error) {
	var records [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}
