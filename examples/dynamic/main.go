package main

import (
	"context"
	"fmt"

	"github.com/dreadster3/gohealth/pkg/healthcheck"
)

func main() {
	service := healthcheck.NewHealthcheckService()

	service.Register("section1", func(ctx context.Context, executor healthcheck.HealthcheckTaskExecutor) healthcheck.HealthcheckStatus {
		executor.Register("check1", func(ctx context.Context, executor healthcheck.HealthcheckTaskExecutor) healthcheck.HealthcheckStatus {
			return healthcheck.HealthcheckStatusHealthy
		})

		executor.Register("check2", func(ctx context.Context, executor healthcheck.HealthcheckTaskExecutor) healthcheck.HealthcheckStatus {
			return healthcheck.HealthcheckStatusHealthy
		})

		return healthcheck.HealthcheckStatusHealthy
	})

	report := service.Run(context.Background())

	// Output: map[section1:healthy section1.check1:healthy section1.check2:healthy]
	fmt.Println(report)
}
