package main

import (
	"github.com/NandoSchlemper/rfm-erp/frontend/views"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/csrf"
)

func main() {
	app := fiber.New()
	app.Use(csrf.New())
	app.Get("/", views.HandleHello)
	app.Listen(":4000")
}
