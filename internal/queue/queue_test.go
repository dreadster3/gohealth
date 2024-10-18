package queue_test

import (
	"testing"
	"time"

	"github.com/dreadster3/gohealth/internal/queue"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := queue.NewWorkerQueue[int](100)

	queue.Enqueue(10)
	queue.Enqueue(20)

	assert.Equal(t, 10, queue.Dequeue())
	assert.Equal(t, 20, queue.Dequeue())
}

func TestQueueBlocking(t *testing.T) {
	t.Parallel()

	queue := queue.NewWorkerQueue[int](100)

	queue.Enqueue(10)

	done := make(chan struct{})

	go func() {
		queue.Wait()
		close(done)
	}()

	select {
	case <-done:
		t.Fatal("should not be done")
	case <-time.After(1 * time.Second):
		return
	}
}

func TestQueueDone(t *testing.T) {
	t.Parallel()

	queue := queue.NewWorkerQueue[int](100)

	queue.Enqueue(10)
	queue.Enqueue(20)

	assert.Equal(t, 10, queue.Dequeue())
	queue.Done()

	done := make(chan struct{})

	go func() {
		queue.Wait()
		close(done)
	}()

	select {
	case <-time.After(1 * time.Second):
		t.Fatal("did not finish")
	case <-done:
		return
	}
}
