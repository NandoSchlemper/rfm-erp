package views

import (
	"log"

	darwin "github.com/NandoSchlemper/rfm-erp/frontend/external/rfm"
	"github.com/NandoSchlemper/rfm-erp/frontend/models"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func HandleForm(c fiber.Ctx) error {
	component := Form()
	handler := adaptor.HTTPHandler(templ.Handler(component))
	return handler(c)
}

func ProcessHandler(c fiber.Ctx) error {
	initialDate := c.FormValue("initial_date")
	finalDate := c.FormValue("final_date")

	log.Print("=== DEBUG FORMS ===")
	log.Printf("Data Inicial: %s", initialDate)
	log.Printf("Data Final: %s", finalDate)

	formData := models.DarwinTrechosRequest{
		Initial_date: initialDate,
		Final_date:   finalDate,
	}

	apiData, err := darwin.GetTrechosData(formData)
	if err != nil {
		errorComponent := Erro("Erro na API: " + err.Error())
		handler := adaptor.HTTPHandler(templ.Handler(errorComponent))
		return handler(c)
	}

	component := Home(apiData)
	handler := adaptor.HTTPHandler(templ.Handler(component))
	return handler(c)
}
