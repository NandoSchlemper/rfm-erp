package main

import (
	"fmt"
	"log"
	"rfmtransportes/external/darwin"
	"rfmtransportes/external/wrsat"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.prod")
	if err != nil {
		fmt.Printf("Erro ao carregar env: %v\n", err)
	}

	fmt.Println("Starting app...")
	app := fiber.New(
		fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		},
	)

	app.Get("/", func(c fiber.Ctx) error {
		data, err := wrsat.GetActualPositions()
		if err != nil {
			fmt.Println("Deu problema...")
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(data)
	})

	app.Get("/darwin/trechos", func(c fiber.Ctx) error {
		data, err := darwin.GetTrechosPercorridos(
			"2026-01-15 00:00",
			"2026-01-15 23:59",
		)

		if err != nil {
			fmt.Printf("Deu problema...\n%v\n", err)
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(data)
	})

	log.Fatal(app.Listen(":3000"))
}
