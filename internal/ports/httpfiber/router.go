package httpfiber

import "github.com/gofiber/fiber/v3"

func GetRouter() *fiber.App {

	app := fiber.New()

	app.Post("/", createShortUrlHandler)

	return app
}
