package main

import (
	"fmt"
	"log"
	"rfmtransportes/external/darwin"

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
	fmt.Println("Starting Repositories...")
	fmt.Println("Starting Services...")
	darwinService := darwin.NewService()

	app := fiber.New(
		fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		},
	)

	app.Get("/darwin/trechos", darwin.GetVehiclesKMByData(darwinService))

	log.Fatal(app.Listen(":3000"))
}
