package healthcheck_test

import (
	"context"
	"testing"
	"time"

	"github.com/dreadster3/gohealth/pkg/healthcheck"
	"github.com/stretchr/testify/assert"
)

func TestServiceRegister(t *testing.T) {
	service := healthcheck.NewHealthcheckService()

	service.Register("test", func(ctx context.Context, executor healthcheck.HealthcheckTaskExecutor) healthcheck.HealthcheckStatus {
		return healthcheck.HealthcheckStatusHealthy
	})

	expected := healthcheck.NewHealthcheckReport()
	expected.Set("test", healthcheck.HealthcheckStatusHealthy)

	actual := service.Run(context.Background())

	assert.Equal(t, expected, actual)
}

func TestServiceDynamicRegister(t *testing.T) {
	service := healthcheck.NewHealthcheckService()

	service.Register("test", func(ctx context.Context, executor healthcheck.HealthcheckTaskExecutor) healthcheck.HealthcheckStatus {
		executor.Register("subtest", func(ctx context.Context, executor healthcheck.HealthcheckTaskExecutor) healthcheck.HealthcheckStatus {
			return healthcheck.HealthcheckStatusHealthy
		})
		return healthcheck.HealthcheckStatusHealthy
	})

	expected := healthcheck.NewHealthcheckReport()
	expected.Set("test", healthcheck.HealthcheckStatusHealthy)
	expected.Set("test.subtest", healthcheck.HealthcheckStatusHealthy)

	actual := service.Run(context.Background())

	assert.Equal(t, expected, actual)
}

func TestServiceRun(t *testing.T) {
	t.Parallel()

	service := healthcheck.NewHealthcheckService()
	done := make(chan struct{})
	done2 := make(chan struct{})

	service.Register("test", func(ctx context.Context, executor healthcheck.HealthcheckTaskExecutor) healthcheck.HealthcheckStatus {
		defer close(done)

		executor.Register("subtest", func(ctx context.Context, executor healthcheck.HealthcheckTaskExecutor) healthcheck.HealthcheckStatus {
			defer close(done2)
			return healthcheck.HealthcheckStatusHealthy
		})

		return healthcheck.HealthcheckStatusHealthy
	})

	go func() {
		service.Run(context.Background())
	}()

	select {
	case <-done:
		break
	case <-time.After(1 * time.Second):
		t.Fatal("timeout")
	}

	select {
	case <-done2:
		break
	case <-time.After(1 * time.Second):
		t.Fatal("timeout")
	}
}
