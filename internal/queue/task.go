package queue

type TaskExecutor[T any] struct {
	name  string
	queue *WorkerQueue[T]
}

func NewTaskExecutor[T any](name string, queue *WorkerQueue[T]) TaskExecutor[T] {
	return TaskExecutor[T]{
		name:  name,
		queue: queue,
	}
}

func (e TaskExecutor[T]) TaskName() string {
	return e.name
}

func (e TaskExecutor[T]) Enqueue(job T) {
	e.queue.Enqueue(job)
}

func (e TaskExecutor[T]) Done() {
	e.queue.Done()
}
