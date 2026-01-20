package views

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func HandleHello(c fiber.Ctx) error {
	component := ShowRFMData()
	handler := adaptor.HTTPHandler(templ.Handler(component))

	return handler(c)
}
