package inmemory

import (
	"errors"
	"sync"
)

type InMemoryRepository struct {
	store *InMemoryStore
}

// implementation InMemoryStore for sync.Map[string]string
type InMemoryStore struct {
	mp sync.Map
}

// put value
func (st *InMemoryStore) Insert(key string, value string) {
	st.mp.Store(key, value)
}

func (r *InMemoryStore) Get(key string) (string, bool) {
	val, ok := r.mp.Load(key)
	if !ok {
		return "", false
	}
	return val.(string), true
}

func (im *InMemoryRepository) Save(url string, id string) error {
	im.store.Insert(id, url)
	return nil
}

func (im *InMemoryRepository) Get(id string) (string, error) {
	real_url, ok := im.store.Get(id)
	if !ok {
		return "", errors.New("no found url")
	}

	return real_url, nil
}
