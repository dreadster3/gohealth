package concurrent_map

import (
	"fmt"
	"sync"
)

type ConcurrentMap[K comparable, V any] struct {
	m  map[K]V
	mu sync.RWMutex
}

func NewConcurrentMap[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{
		m: make(map[K]V),
	}
}

func (m *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, ok := m.m[key]
	return value, ok
}

func (m *ConcurrentMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.m[key] = value
}

func (m *ConcurrentMap[K, V]) Delete(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.m, key)
}

func (m *ConcurrentMap[K, V]) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.m)
}

// Items returns a copy of the internal map.
func (m *ConcurrentMap[K, V]) Iter() func(func(K, V) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return func(yield func(key K, value V) bool) {
		for k, v := range m.m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (m *ConcurrentMap[K, V]) String() string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return fmt.Sprintf("%v", m.m)
}
