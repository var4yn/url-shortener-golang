package httpfiber

import (
	"url-shortener/internal/depend"

	"github.com/gofiber/fiber/v3"
)

func GetRouter(appState *depend.AppState) *fiber.App {

	app := fiber.New()

	app.State().Set("appState", appState)

	// add state
	app.Use(func(c fiber.Ctx) error {
		c.Locals("state", appState)
		return c.Next()
	})
	app.Post("/", createShortUrlHandler)
	app.Get("/:id", getRealUrlHandler)

	return app
}
