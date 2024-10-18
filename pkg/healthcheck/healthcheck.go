package healthcheck

import "context"

type HealthcheckFn func(context.Context, HealthcheckTaskExecutor) HealthcheckStatus

type Healthcheck struct {
	name string
	fn   HealthcheckFn
}

func NewHealthcheck(name string, fn HealthcheckFn) *Healthcheck {
	return &Healthcheck{
		name: name,
		fn:   fn,
	}
}

func (h *Healthcheck) Run(ctx context.Context, executor HealthcheckTaskExecutor) HealthcheckStatus {
	defer executor.done()
	return h.fn(ctx, executor)
}
