package httpfiber

import (
	"url-shortener/internal/depend"

	"github.com/gofiber/fiber/v3"
)

type RealUrlResponse struct {
	URL string `json:"url"`
}

func getRealUrlHandler(c fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	state, ok := c.Locals("state").(*depend.AppState)

	if !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	url := ""

	res := RealUrlResponse{URL: url}
	return c.JSON(res)
}
