package main

import (
	"github.com/AhmadIkbalDjaya/go-simple-pos/config"
	"github.com/AhmadIkbalDjaya/go-simple-pos/exception"
	"github.com/AhmadIkbalDjaya/go-simple-pos/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})
	
	routes.SetUpRoutes(app)
	app.Listen(":"+config.Config.AppPort)
}