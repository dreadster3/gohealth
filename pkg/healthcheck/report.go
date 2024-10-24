package healthcheck

import (
	"encoding/json"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/dreadster3/gohealth/internal/concurrent_map"
)

type HealthcheckReport struct {
	*concurrent_map.ConcurrentMap[string, HealthcheckStatus]
}

// JSONHealthcheckReport is a struct that represents the JSON output of a healthcheck report when JSON encoded.
type JSONHealthcheckReport struct {
	Status   HealthcheckStatus            `json:"status"`
	Services map[string]HealthcheckStatus `json:"services"`
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
	result := map[string]bool{}
	for key := range r.Iter() {
		splits := strings.Split(key, ".")

		if len(splits) > 1 {
			result[splits[0]] = true
		}
	}

	return slices.Collect(maps.Keys(result))
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

func (r HealthcheckReport) MarshalJSON() ([]byte, error) {
	result := struct {
		Status   HealthcheckStatus                                        `json:"status"`
		Services *concurrent_map.ConcurrentMap[string, HealthcheckStatus] `json:"services"`
	}{
		Status:   r.Status(),
		Services: r.ConcurrentMap,
	}

	return json.Marshal(result)
}
