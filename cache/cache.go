package cache

import (
	"context"
	"fmt"
	"sync"

	"github.com/fogfish/skiplist"
	"github.com/fogfish/skiplist/ord"
)

type Cache[K ord.Comparable, V any] struct {
	lock  *sync.Mutex
	store *skiplist.SkipList[K, V]
}

func New[K ord.Comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		lock:  &sync.Mutex{},
		store: skiplist.New[K, V](ord.Type[K]()),
	}
}

func (cache *Cache[K, V]) Get(_ context.Context, key K) (V, error) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	val, has := skiplist.Lookup(cache.store, key)
	if !has {
		return val, errNotFound(fmt.Sprintf("%v", key))
	}

	return val, nil
}

func (cache *Cache[K, V]) Seq(_ context.Context, afterKey K, size int) ([]V, error) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	seq := make([]V, 0)

	_, tail := skiplist.Split(cache.store, afterKey)
	if tail == nil {
		return seq, nil
	}

	for tail.Next() {
		_, val := tail.Head()
		seq = append(seq, val)

		if len(seq) == size {
			return seq, nil
		}
	}

	return seq, nil
}

func (cache *Cache[K, V]) Set(_ context.Context, key K, val V) error {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	skiplist.Put(cache.store, key, val)
	return nil
}

type errNotFound string

func (key errNotFound) Error() string    { return string(key) }
func (key errNotFound) NotFound() string { return string(key) }
