package queue_test

import (
	"testing"
	"time"

	"github.com/dreadster3/gohealth/internal/queue"
	"github.com/stretchr/testify/assert"
)

func TestTaskExecutor(t *testing.T) {
	workerQueue := queue.NewWorkerQueue[int](100)
	executor := queue.NewTaskExecutor("test", workerQueue)

	assert.Equal(t, "test", executor.TaskName())

	executor.Enqueue(1)
	executor.Enqueue(2)

	assert.Equal(t, 1, workerQueue.Dequeue())
	assert.Equal(t, 2, workerQueue.Dequeue())

	done := make(chan struct{})

	go func() {
		workerQueue.Wait()
		close(done)
	}()

	select {
	case <-done:
		return
	case <-time.After(1 * time.Second):
		t.Fatal("timeout")
	}
}

func TestTaskExecutorBlocking(t *testing.T) {
	workerQueue := queue.NewWorkerQueue[int](100)
	executor := queue.NewTaskExecutor("test", workerQueue)

	assert.Equal(t, "test", executor.TaskName())

	executor.Enqueue(1)
	executor.Enqueue(2)

	assert.Equal(t, 1, workerQueue.Dequeue())

	done := make(chan struct{})

	go func() {
		workerQueue.Wait()
		close(done)
	}()

	select {
	case <-done:
		t.Fatal("should not be done")
	case <-time.After(1 * time.Second):
		return
	}
}
