package query

// interface for repo
type GetRealUrlRepository interface {
	Get(id string) (string, error)
}

// query struct
type GetRealUrlQuery struct {
	repo GetRealUrlRepository
}

// constructor
func NewGetRealUrlQuery(repo GetRealUrlRepository) GetRealUrlQuery {
	return GetRealUrlQuery{repo}
}

// execute method
func (gr *GetRealUrlQuery) Execute(id string) (string, error) {
	return gr.repo.Get(id)
}
