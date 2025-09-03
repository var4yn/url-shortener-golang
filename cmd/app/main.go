package main

import (
	"log"
	"url-shortener/internal/ports/httpfiber"
)

func main() {
	app := httpfiber.GetRouter()

	log.Fatal(app.Listen(":3000"))
}
