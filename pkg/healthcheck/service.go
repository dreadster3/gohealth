package healthcheck

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

func (s *HealthcheckService) Run(ctx context.Context) HealthcheckReport {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		s.queue.Wait()
		cancel()
	}()

	result := NewHealthcheckReport()
	for {
		select {
		case <-ctx.Done():
			return result
		case h := <-s.queue.Chan():
			go func() {
				status := h.Run(ctx, NewHealthCheckExecutor(h.name, s.queue))
				result.Set(h.name, status)
			}()
		}
	}
}
