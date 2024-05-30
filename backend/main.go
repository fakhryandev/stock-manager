package main

import (
	"log"
	"stock-manager/database"
	"stock-manager/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	app.Use(logger.New())
	app.Use(cors.New())

	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Println("Server running on localhost:8080")

	err := app.Listen(":8080")

	if err != nil {
		log.Fatal(err.Error())
	}
}
