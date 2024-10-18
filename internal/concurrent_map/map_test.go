package concurrent_map_test

import (
	"math/rand/v2"
	"sync"
	"testing"

	"github.com/dreadster3/gohealth/internal/concurrent_map"
	"github.com/stretchr/testify/assert"
)

func TestConcurrentMapIter(t *testing.T) {
	t.Parallel()

	concurrentMap := concurrent_map.NewConcurrentMap[string, int]()

	concurrentMap.Set("key1", 1)
	concurrentMap.Set("key2", 2)

	tmp := map[string]int{}
	for key, value := range concurrentMap.Iter() {
		tmp[key] = value
	}

	assert.Equal(t, map[string]int{"key1": 1, "key2": 2}, tmp)
}

func TestConcurrentMapSet(t *testing.T) {
	t.Parallel()

	concurrentMap := concurrent_map.NewConcurrentMap[string, int]()

	concurrentMap.Set("key", 1)

	value, ok := concurrentMap.Get("key")
	assert.True(t, ok)
	assert.Equal(t, 1, value)
}

func TestConcurrentMapDelete(t *testing.T) {
	t.Parallel()

	concurrentMap := concurrent_map.NewConcurrentMap[string, int]()

	concurrentMap.Set("key", 1)
	concurrentMap.Delete("key")

	_, ok := concurrentMap.Get("key")
	assert.False(t, ok)
}

func TestConcurrentMapLen(t *testing.T) {
	t.Parallel()

	concurrentMap := concurrent_map.NewConcurrentMap[string, int]()

	concurrentMap.Set("key1", 1)
	concurrentMap.Set("key2", 2)

	assert.Equal(t, 2, concurrentMap.Len())
}

func TestConcurrentMapConcurrentWrites(t *testing.T) {
	t.Parallel()

	concurrentMap := concurrent_map.NewConcurrentMap[string, int]()

	go func() {
		for i := 0; i < 1000; i++ {
			concurrentMap.Set("key", i)
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			concurrentMap.Set("key", i)
		}
	}()

	for i := 0; i < 1000; i++ {
		concurrentMap.Set("key", i)
	}

	assert.Equal(t, 1, concurrentMap.Len())
}

func TestConcurrentMapConcurrentReadsAndWrites(t *testing.T) {
	t.Parallel()

	concurrentMap := concurrent_map.NewConcurrentMap[string, int]()

	for i := 0; i < 1000; i++ {
		concurrentMap.Set("key", i)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			concurrentMap.Set("key", rand.Int())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			_, ok := concurrentMap.Get("key")
			assert.True(t, ok)
		}
	}()

	wg.Wait()

	assert.Equal(t, 1, concurrentMap.Len())
}
