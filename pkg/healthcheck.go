package pkg

type HealthcheckFn func(HealthcheckContext) HealthcheckStatus

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

func (h *Healthcheck) Run(ctx HealthcheckContext) HealthcheckStatus {
	defer ctx.queue.waitGroup.Done()
	return h.fn(ctx)
}
