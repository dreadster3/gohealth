# gohealth

Health check library for Go applications. Supports dynamic registrations of healthchecks.

## Usage

See [here](https://github.com/dreadster3/gohealth/blob/main/examples) for more examples.

```go
service := healthcheck.NewHealthcheckService()

service.Register("healthcheck1", func(ctx context.Context, executor healthcheck.HealthcheckTaskExecutor) healthcheck.HealthcheckStatus {
	return healthcheck.HealthcheckStatusHealthy
})

service.Register("healthcheck2", func(ctx context.Context, executor healthcheck.HealthcheckTaskExecutor) healthcheck.HealthcheckStatus {
	return healthcheck.HealthcheckStatusHealthy
})

report := service.Run(context.Background())

// Output: map[healthcheck1:healthy healthcheck2:healthy]
fmt.Println(report)
```
