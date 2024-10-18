package healthcheck

import (
	"fmt"
	"strings"

	"github.com/dreadster3/gohealth/internal/concurrent_map"
)

type HealthcheckReport struct {
	*concurrent_map.ConcurrentMap[string, HealthcheckStatus]
}

func NewHealthcheckReport() HealthcheckReport {
	return HealthcheckReport{
		concurrent_map.NewConcurrentMap[string, HealthcheckStatus](),
	}
}

func (r HealthcheckReport) Status() HealthcheckStatus {
	status := HealthcheckStatusHealthy
	for _, value := range r.Iter() {
		if value == HealthcheckStatusUnhealthy {
			return HealthcheckStatusUnhealthy
		}

		if value == HealthcheckStatusDegraded {
			status = HealthcheckStatusDegraded
		}
	}
	return status
}

func (r HealthcheckReport) GetSectionsName() []string {
	result := []string{}
	for key := range r.Iter() {
		splits := strings.Split(key, ".")

		if len(splits) > 1 {
			result = append(result, splits[0])
		}
	}

	return result
}

func (r HealthcheckReport) GetSection(section string) HealthcheckReport {
	result := NewHealthcheckReport()
	for key, value := range r.Iter() {
		if strings.HasPrefix(key, fmt.Sprintf("%s.", section)) {
			key := strings.TrimPrefix(key, fmt.Sprintf("%s.", section))
			result.Set(key, value)
		}
	}

	return result
}

func (r HealthcheckReport) GetIndividualCheckStatus() HealthcheckReport {
	result := NewHealthcheckReport()
	for key, value := range r.Iter() {
		if !strings.Contains(key, ".") {
			result.Set(key, value)
		}
	}

	return result
}
