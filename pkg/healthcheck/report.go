package healthcheck

import (
	"fmt"
	"strings"
)

type HealthcheckReport map[string]HealthcheckStatus

func NewHealthcheckReport() HealthcheckReport {
	return HealthcheckReport{}
}

func (r HealthcheckReport) Status() HealthcheckStatus {
	status := HealthcheckStatusHealthy
	for _, value := range r {
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
	for key := range r {
		splits := strings.Split(key, ".")

		if len(splits) > 1 {
			result = append(result, splits[0])
		}
	}

	return result
}

func (r HealthcheckReport) GetSection(section string) HealthcheckReport {
	result := HealthcheckReport{}
	for key, value := range r {
		if strings.HasPrefix(key, fmt.Sprintf("%s.", section)) {
			key := strings.TrimPrefix(key, fmt.Sprintf("%s.", section))
			result[key] = value
		}
	}

	return result
}

func (r HealthcheckReport) GetIndividualCheckStatus() HealthcheckReport {
	result := HealthcheckReport{}
	for key, value := range r {
		if !strings.Contains(key, ".") {
			result[key] = value
		}
	}

	return result
}
