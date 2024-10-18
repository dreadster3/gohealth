package pkg

type HealthcheckStatus string

const (
	HealthcheckStatusHealthy   HealthcheckStatus = "healthy" 
	HealthcheckStatusUnhealthy HealthcheckStatus = "unhealthy"
	HealthcheckStatusDegraded  HealthcheckStatus = "degraded"
)

