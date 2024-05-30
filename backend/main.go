package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	log.Println("Server running on localhost:8080")
	app.Listen("8080")
}
