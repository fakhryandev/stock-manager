package router

import (
	"stock-manager/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupItemRoutes(router fiber.Router) {
	items := router.Group("/item")

	items.Get("/", handlers.GetItems)
	items.Get("/:code", handlers.GetItem)
	items.Post("/", handlers.CreateItem)
	items.Put("/:code", handlers.UpdateItem)
	items.Delete("/:code", handlers.DeleteItem)
	items.Patch("/:code/increase", handlers.IncreaseItem)
	items.Patch("/:code/decrease", handlers.DecreaseItem)
}
