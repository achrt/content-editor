package main

import (
	"content-editor/cmd/editor/application"
	"content-editor/cmd/editor/handler"
	"github.com/labstack/gommon/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	router := fiber.New(defaultFiberConfig)

	router.Use(
		recover.New(),
	)

	app, err := application.New(router)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.New(app)
	h.Route(router)

	log.Fatal(app.Run())
}

const fiberBufferSize = 1024 * 16

var defaultFiberConfig = fiber.Config{
	DisableStartupMessage: true,
	ErrorHandler:          errorHandler,
	ReadBufferSize:        fiberBufferSize,
	WriteBufferSize:       fiberBufferSize,
}

var errorHandler = func(c *fiber.Ctx, err error) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(fiber.StatusInternalServerError).JSON(err)
}
