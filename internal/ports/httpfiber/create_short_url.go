package httpfiber

import (
	"url-shortener/internal/depend"

	"github.com/gofiber/fiber/v3"
)

type ShortUrlRequest struct {
	URL string `json:"url"`
}

type ShortUrlResponse struct {
	ID string `json:"id"`
}

func createShortUrlHandler(c fiber.Ctx) error {
	var req ShortUrlRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid request"})
	}

	state, ok := fiber.GetState[*depend.AppState](c.App().State(), "appState")

	if !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	id, err := state.CreateShortUrlCommand.Execute(req.URL)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	res := ShortUrlResponse{ID: id}
	return c.JSON(res)
}
