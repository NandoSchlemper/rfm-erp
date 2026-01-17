package darwin

import "github.com/gofiber/fiber/v3"

func GetVehiclesKMByData(service Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		var reqPayload DarwinAPITrechosPayload

		err := c.Bind().Body(reqPayload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadGateway, "Do the request right. Motherfucker")
		}

		response, err := service.GetVehiclesKM(reqPayload)
		if err != nil || len(response) == 0 {
			return fiber.NewError(fiber.StatusBadGateway, "Erro ao processar os dados da API.")
		}

		return c.JSON(response)
	}
}
