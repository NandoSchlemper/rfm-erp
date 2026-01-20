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

	priorityPlacas := []string{
		"RDX-4D10",
		"FVU-0537",
		"SDU-7E80",
		"SCM-2A20",
		"SCM-4A40",
		"SCM-6A60",
		"SFV-3E38",
		"SFU-1E58",
		"RSB-7A50",
		"SFU-1E57",
		"RSC-6C92",
		"RSF-1I49",
		"RSD-7B09",
		"RSB-6J00",
		"SFU-2A17",
		"RHY-2I31",
		"SGA-2I15",
		"BDY-9G00",
		"SFW-5G99",
	}

	sortedData := darwin.SortPlacasByPriorityAndKM(apiData, priorityPlacas)

	component := Home(sortedData)
	handler := adaptor.HTTPHandler(templ.Handler(component))

	return handler(c)
}
