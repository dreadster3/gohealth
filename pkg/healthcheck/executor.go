package healthcheck

import (
	"fmt"

	"github.com/dreadster3/gohealth/internal/queue"
)

// HealthcheckQueue is an alias for the WorkerQueue type for Healthchecks.
type HealthcheckQueue = queue.WorkerQueue[*Healthcheck]

func NewHealthcheckQueue() *HealthcheckQueue {
	return queue.NewWorkerQueue[*Healthcheck](100)
}

// HealthcheckTaskExecutor is an alias for the TaskExecutor type for Healthchecks.
type HealthcheckTaskExecutor queue.TaskExecutor[*Healthcheck]

func NewHealthCheckExecutor(name string, q *HealthcheckQueue) HealthcheckTaskExecutor {
	return (HealthcheckTaskExecutor)(queue.NewTaskExecutor(name, q))
}

func (e HealthcheckTaskExecutor) TaskName() string {
	return (queue.TaskExecutor[*Healthcheck])(e).TaskName()
}

func (e HealthcheckTaskExecutor) Register(name string, fn HealthcheckFn) {
	(queue.TaskExecutor[*Healthcheck])(e).Enqueue(NewHealthcheck(fmt.Sprintf("%s.%s", e.TaskName(), name), fn))
}

func (e HealthcheckTaskExecutor) done() {
	(queue.TaskExecutor[*Healthcheck])(e).Done()
}
