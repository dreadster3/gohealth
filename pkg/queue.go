package pkg

import "sync"

type HealthcheckQueue struct {
	queue     chan *Healthcheck
	waitGroup *sync.WaitGroup
}

func NewHealthcheckQueue() *HealthcheckQueue {
	return &HealthcheckQueue{
		queue:     make(chan *Healthcheck, 100),
		waitGroup: &sync.WaitGroup{},
	}
}

func (q *HealthcheckQueue) Enqueue(h *Healthcheck) {
	q.waitGroup.Add(1)
	q.queue <- h
}

func (q *HealthcheckQueue) Dequeue() *Healthcheck {
	defer q.waitGroup.Done()
	return <-q.queue
}
