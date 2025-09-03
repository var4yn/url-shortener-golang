package command

import (
	"net/url"
	"url-shortener/internal/generator"
)

type CreateShortUrlRepository interface {
	Save(url string, id string) error
}

type CreateShortUrlCommand struct {
	repo      CreateShortUrlRepository
	generator generator.IdGenerator
}

func NewCreateShortUrlCommand(repo CreateShortUrlRepository, generator generator.IdGenerator) CreateShortUrlCommand {
	return CreateShortUrlCommand{
		repo:      repo,
		generator: generator,
	}
}

func (cs *CreateShortUrlCommand) execute(real_url string) (string, error) {
	// url validate
	u, err := url.ParseRequestURI(real_url)
	if err != nil {
		return "", err
	}

	id := cs.generator.Generate()

	err = cs.repo.Save(u.String(), id)

	return id, err
}
