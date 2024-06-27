package server

import (
	"log"
	"module/internal/models"

	"github.com/gofiber/fiber/v2"
)

func ServerStart(port models.ServerConfig) *fiber.App {

	app := fiber.New()

	handlers(app)

	log.Fatal(app.Listen(port.Port))

	return app
}

func handlers(instance *fiber.App) {

	instance.Get("/book", zaglushka)
	instance.Post("/book", zaglushka)

	instance.Get("/client", zaglushka)
	instance.Post("/client", zaglushka)

}

func zaglushka(c *fiber.Ctx) error {
	return models.ResponseGood()
}
