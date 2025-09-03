package query

// interface for repo
type GetRealUrlRepository interface {
	get(id string) (string, error)
}

// query struct
type GetRealUrlQuery struct {
	repo GetRealUrlRepository
}

// constructor
func NewGetRealUrl(repo GetRealUrlRepository) GetRealUrlQuery {
	return GetRealUrlQuery{repo}
}

// execute method
func (gr *GetRealUrlQuery) execute(id string) (string, error) {
	return gr.repo.get(id)
}
