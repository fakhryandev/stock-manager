package router

import (
	"stock-manager/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupItemRoutes(router fiber.Router) {
	items := router.Group("/item")

	items.Get("/", handlers.GetItems)
	items.Post("/", handlers.CreateItem)
}
