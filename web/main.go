package main

import (
	"github.com/NandoSchlemper/rfm-erp/frontend/views"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/csrf"
)

func main() {
	app := fiber.New()
	app.Use(csrf.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	app.Get("/", views.HandleForm)
	app.Get("/process", views.ProcessHandler)
	app.Listen(":4000")
}
