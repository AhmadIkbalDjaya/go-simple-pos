package main

import (
	categorycontroller "github.com/AhmadIkbalDjaya/go-simple-pos/controller/category_controller"
	"github.com/AhmadIkbalDjaya/go-simple-pos/exception"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})
	
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
	
	app.Listen(":3000")
}