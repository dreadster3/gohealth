package pkg

import (
	"context"
)

type HealthcheckService struct {
	queue *HealthcheckQueue
}

func NewHealthcheckService() *HealthcheckService {
	return &HealthcheckService{
		queue: NewHealthcheckQueue(),
	}
}

func (s *HealthcheckService) Register(name string, fn HealthcheckFn) {
	s.queue.Enqueue(NewHealthcheck(name, fn))
}

func (s *HealthcheckService) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		s.queue.waitGroup.Wait()
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case h := <-s.queue.queue:
			go h.Run(NewHealthcheckContext(ctx, h.name, s.queue))
		}
	}
}
