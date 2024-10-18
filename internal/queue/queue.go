package queue

import "sync"

type WorkerQueue[T any] struct {
	queue     chan T
	waitGroup sync.WaitGroup
}

func NewWorkerQueue[T any](bufferSize int) *WorkerQueue[T] {
	return &WorkerQueue[T]{
		queue:     make(chan T, bufferSize),
		waitGroup: sync.WaitGroup{},
	}
}

// Enqueue adds a new element to the queue.
func (q *WorkerQueue[T]) Enqueue(job T) {
	q.waitGroup.Add(1)
	q.queue <- job
}

// Dequeue removes the first element from the queue and returns it.
func (q *WorkerQueue[T]) Dequeue() T {
	defer q.Done()
	return <-q.queue
}

// Wait blocks until all workers have finished processing the queue. That is until the queue is empty.
func (q *WorkerQueue[T]) Wait() {
	q.waitGroup.Wait()
}

// Done decrements the wait group counter by one. That is, it signals that a worker has finished processing a job.
func (q *WorkerQueue[T]) Done() {
	q.waitGroup.Done()
}

// Chan returns the channel of the queue.
func (q *WorkerQueue[T]) Chan() chan T {
	return q.queue
}
