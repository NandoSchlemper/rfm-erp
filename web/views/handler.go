package views

import (
	"github.com/NandoSchlemper/rfm-erp/frontend/api/darwin"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func HandleHello(c fiber.Ctx) error {
	apiData, err := darwin.GetTrechosData()
	if err != nil {
		return c.Status(500).SendString("Erro ao buscar dados: " + err.Error())
	}

	component := Home(apiData)
	handler := adaptor.HTTPHandler(templ.Handler(component))

	return handler(c)
}
