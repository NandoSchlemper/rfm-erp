package views

import (
	darwin "github.com/NandoSchlemper/rfm-erp/frontend/external/rfm"
	"github.com/NandoSchlemper/rfm-erp/frontend/models"
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

func HandleForm(c fiber.Ctx) error {
	component := Form()
	handler := adaptor.HTTPHandler(templ.Handler(component))

	return handler(c)
}

func ProcessHandler(c fiber.Ctx) error {
	var formData models.DarwinTrechosRequest

	if err := c.Bind().Body(&formData); err != nil {
		return c.Status(400).SendString("Erro ao processar dados " + err.Error())
	}

	apiData, err := darwin.GetTrechosData(formData)
	if err != nil {
		return c.Status(400).SendString("Erro no retorno da API" + err.Error())
	}

	component := Home(apiData)
	return component.Render(c.Context(), c.Response().BodyWriter())
}
