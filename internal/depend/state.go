package depend

import (
	"url-shortener/internal/app/command"
	"url-shortener/internal/app/query"
	"url-shortener/internal/generator"
)

type AppState struct {
	CreateShortUrlCommand command.CreateShortUrlCommand
	GetRealUrlQuery       query.GetRealUrlQuery
}

func NewAppState(
	generator generator.IdGenerator,
	create_short_url_repo command.CreateShortUrlRepository,
	get_real_url_repo query.GetRealUrlRepository) AppState {

	return AppState{
		CreateShortUrlCommand: command.NewCreateShortUrlCommand(create_short_url_repo, generator),
		GetRealUrlQuery:       query.NewGetRealUrlQuery(get_real_url_repo),
	}

}
