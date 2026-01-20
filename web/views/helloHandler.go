package views

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func HandleHello(c fiber.Ctx) error {
	helloComponent := Hello("Hi there!")
	handler := adaptor.HTTPHandler(templ.Handler(helloComponent))

	return handler(c)
}
