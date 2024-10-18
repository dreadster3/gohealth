package healthcheck_test

import (
	"sort"
	"testing"

	"github.com/dreadster3/gohealth/pkg/healthcheck"
	"github.com/stretchr/testify/assert"
)

func TestReportGetSectionsName(t *testing.T) {
	report := healthcheck.NewHealthcheckReport()

	report.Set("section1.check1", healthcheck.HealthcheckStatusHealthy)
	report.Set("section2.check1", healthcheck.HealthcheckStatusHealthy)
	report.Set("section3.check1", healthcheck.HealthcheckStatusHealthy)

	expected := []string{"section1", "section2", "section3"}
	actual := report.GetSectionsName()

	sort.Strings(actual)

	assert.Equal(t, expected, actual)
}

func TestReportGetSection(t *testing.T) {
	report := healthcheck.NewHealthcheckReport()

	report.Set("section1.check1", healthcheck.HealthcheckStatusHealthy)
	report.Set("section1.check2", healthcheck.HealthcheckStatusHealthy)
	report.Set("section2.check1", healthcheck.HealthcheckStatusHealthy)
	report.Set("section3.check1", healthcheck.HealthcheckStatusHealthy)

	expected := healthcheck.NewHealthcheckReport()
	expected.Set("check1", healthcheck.HealthcheckStatusHealthy)
	expected.Set("check2", healthcheck.HealthcheckStatusHealthy)

	section := report.GetSection("section1")

	assert.Equal(t, expected, section)
}

func TestReportStatusHealthy(t *testing.T) {
	report := healthcheck.NewHealthcheckReport()

	report.Set("check1", healthcheck.HealthcheckStatusHealthy)
	report.Set("check2", healthcheck.HealthcheckStatusHealthy)
	report.Set("check3", healthcheck.HealthcheckStatusHealthy)

	expected := healthcheck.HealthcheckStatusHealthy
	actual := report.Status()

	assert.Equal(t, expected, actual)
}

func TestReportStatusDegraded(t *testing.T) {
	report := healthcheck.NewHealthcheckReport()

	report.Set("check1", healthcheck.HealthcheckStatusHealthy)
	report.Set("check2", healthcheck.HealthcheckStatusDegraded)
	report.Set("check3", healthcheck.HealthcheckStatusHealthy)

	expected := healthcheck.HealthcheckStatusDegraded
	actual := report.Status()

	assert.Equal(t, expected, actual)
}

func TestReportStatusUnhealthy(t *testing.T) {
	report := healthcheck.NewHealthcheckReport()

	report.Set("check1", healthcheck.HealthcheckStatusHealthy)
	report.Set("check2", healthcheck.HealthcheckStatusUnhealthy)
	report.Set("check3", healthcheck.HealthcheckStatusDegraded)

	expected := healthcheck.HealthcheckStatusUnhealthy
	actual := report.Status()

	assert.Equal(t, expected, actual)
}

func TestReportGetIndividualCheckStatus(t *testing.T) {
	report := healthcheck.NewHealthcheckReport()

	report.Set("check1", healthcheck.HealthcheckStatusHealthy)
	report.Set("check2", healthcheck.HealthcheckStatusUnhealthy)
	report.Set("check3", healthcheck.HealthcheckStatusDegraded)
	report.Set("section1.check1", healthcheck.HealthcheckStatusHealthy)
	report.Set("section1.check2", healthcheck.HealthcheckStatusHealthy)
	report.Set("section1.check3", healthcheck.HealthcheckStatusHealthy)

	expected := healthcheck.NewHealthcheckReport()
	expected.Set("check1", healthcheck.HealthcheckStatusHealthy)
	expected.Set("check2", healthcheck.HealthcheckStatusUnhealthy)
	expected.Set("check3", healthcheck.HealthcheckStatusDegraded)

	actual := report.GetIndividualCheckStatus()

	assert.Equal(t, expected, actual)
}
