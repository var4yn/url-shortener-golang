package httpfiber

import "github.com/gofiber/fiber/v3"

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

	id := "abc123"

	res := ShortUrlResponse{ID: id}
	return c.JSON(res)
}
