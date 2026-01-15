package main

import (
	"fmt"
	"log"
	"rfmtransportes/external/wrsat"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
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

	log.Fatal(app.Listen(":3000"))
}
