package store

import "sync"

type Repository[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

func NewRepository[K comparable, V any]() *Repository[K, V] {
	return &Repository[K, V]{
		data: make(map[K]V),
	}
}

func (r *Repository[K, V]) Set(key K, value V) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[key] = value
}

func (r *Repository[K, V]) Get(key K) (V, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	val, ok := r.data[key]
	return val, ok
}

func (r *Repository[K, V]) All() []V {
	r.mu.RLock()
	defer r.mu.RUnlock()

	res := []V{}
	for _, v := range r.data {
		res = append(res, v)
	}
	return res
}

func (r *Repository[K, V]) Delete(key K) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.data, key)
}
