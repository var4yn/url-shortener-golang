package main

import (
	"log"
	"url-shortener/internal/adapters/inmemory"
	"url-shortener/internal/depend"
	"url-shortener/internal/generator"
	"url-shortener/internal/ports/httpfiber"
)

func main() {

	generator := generator.NanoidGenerator{}
	var store inmemory.InMemoryStore

	create_short_url_repo := inmemory.NewInMemoryRepository(&store)
	get_real_url_repo := inmemory.NewInMemoryRepository(&store)

	state := depend.NewAppState(generator, create_short_url_repo, get_real_url_repo)

	app := httpfiber.GetRouter(&state)

	log.Fatal(app.Listen(":3000"))
}
