package routes

import (
	categorycontroller "github.com/AhmadIkbalDjaya/go-simple-pos/controller/category_controller"
	productcontroller "github.com/AhmadIkbalDjaya/go-simple-pos/controller/product_controller"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	api := app.Group("/api")
	
	category := api.Group("/categories")
	category.Get("/", categorycontroller.Index)
	category.Get("/:categoryId", categorycontroller.Show)
	category.Post("/", categorycontroller.Create)
	category.Put("/:categoryId", categorycontroller.Update)
	category.Delete("/:categoryId", categorycontroller.Delete)
	
	product := api.Group("/products")
	product.Get("/", productcontroller.Index)
	product.Get("/:productId", productcontroller.Show)	
	product.Post("/", productcontroller.Create)
	product.Put("/:productId", productcontroller.Update)
	product.Delete("/:productId", productcontroller.Delete)
}