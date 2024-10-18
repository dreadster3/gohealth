package healthcheck_test

import (
	"sort"
	"testing"

	"github.com/dreadster3/gohealth/pkg/healthcheck"
	"github.com/stretchr/testify/assert"
)

func TestReportGetSectionsName(t *testing.T) {
	report := healthcheck.NewHealthcheckReport()

	report["section1.check1"] = healthcheck.HealthcheckStatusHealthy
	report["section2.check1"] = healthcheck.HealthcheckStatusHealthy
	report["section3.check1"] = healthcheck.HealthcheckStatusHealthy

	expected := []string{"section1", "section2", "section3"}
	actual := report.GetSectionsName()

	sort.Strings(actual)

	assert.Equal(t, expected, actual)
}

func TestReportGetSection(t *testing.T) {
	report := healthcheck.NewHealthcheckReport()

	report["section1.check1"] = healthcheck.HealthcheckStatusHealthy
	report["section1.check2"] = healthcheck.HealthcheckStatusHealthy
	report["section2.check1"] = healthcheck.HealthcheckStatusHealthy
	report["section3.check1"] = healthcheck.HealthcheckStatusHealthy

	expected := healthcheck.NewHealthcheckReport()
	expected["check1"] = healthcheck.HealthcheckStatusHealthy
	expected["check2"] = healthcheck.HealthcheckStatusHealthy

	section := report.GetSection("section1")

	assert.Equal(t, expected, section)
}

func TestReportStatusHealthy(t *testing.T) {
	report := healthcheck.NewHealthcheckReport()

	report["check1"] = healthcheck.HealthcheckStatusHealthy
	report["check2"] = healthcheck.HealthcheckStatusHealthy
	report["check3"] = healthcheck.HealthcheckStatusHealthy

	expected := healthcheck.HealthcheckStatusHealthy
	actual := report.Status()

	assert.Equal(t, expected, actual)
}

func TestReportStatusDegraded(t *testing.T) {
	report := healthcheck.NewHealthcheckReport()

	report["check1"] = healthcheck.HealthcheckStatusHealthy
	report["check2"] = healthcheck.HealthcheckStatusDegraded
	report["check3"] = healthcheck.HealthcheckStatusHealthy

	expected := healthcheck.HealthcheckStatusDegraded
	actual := report.Status()

	assert.Equal(t, expected, actual)
}

func TestReportStatusUnhealthy(t *testing.T) {
	report := healthcheck.NewHealthcheckReport()

	report["check1"] = healthcheck.HealthcheckStatusHealthy
	report["check2"] = healthcheck.HealthcheckStatusUnhealthy
	report["check3"] = healthcheck.HealthcheckStatusDegraded

	expected := healthcheck.HealthcheckStatusUnhealthy
	actual := report.Status()

	assert.Equal(t, expected, actual)
}

func TestReportGetIndividualCheckStatus(t *testing.T) {
	report := healthcheck.NewHealthcheckReport()

	report["check1"] = healthcheck.HealthcheckStatusHealthy
	report["check2"] = healthcheck.HealthcheckStatusUnhealthy
	report["check3"] = healthcheck.HealthcheckStatusDegraded
	report["section1.check1"] = healthcheck.HealthcheckStatusHealthy
	report["section1.check2"] = healthcheck.HealthcheckStatusHealthy
	report["section1.check3"] = healthcheck.HealthcheckStatusHealthy

	expected := healthcheck.HealthcheckReport{
		"check1": healthcheck.HealthcheckStatusHealthy,
		"check2": healthcheck.HealthcheckStatusUnhealthy,
		"check3": healthcheck.HealthcheckStatusDegraded,
	}

	actual := report.GetIndividualCheckStatus()

	assert.Equal(t, expected, actual)
}
