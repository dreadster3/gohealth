package pkg

import "context"

type HealthcheckContext struct {
	name    string
	context context.Context
	queue   *HealthcheckQueue
}

func NewHealthcheckContext(ctx context.Context, name string, queue *HealthcheckQueue) HealthcheckContext {
	return HealthcheckContext{
		context: ctx,
		name:    name,
		queue:   queue,
	}
}

func (c HealthcheckContext) Name() string {
	return c.name
}

func (c HealthcheckContext) AddCheck(name string, fn HealthcheckFn) {
	c.queue.Enqueue(NewHealthcheck(name, fn))
}
