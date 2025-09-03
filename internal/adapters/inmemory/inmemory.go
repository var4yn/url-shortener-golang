package inmemory

import "sync"

type InMemoryRepository struct {
	store InMemoryStore
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
