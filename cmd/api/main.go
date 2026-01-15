package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	fmt.Println("Starting app...")
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello there!")
	})

	log.Fatal(app.Listen(":3000"))
}
