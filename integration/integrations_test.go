package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
)

func init() {
	client = initTest()
}

func TestIntegrationsPDCreateAndDelete(t *testing.T) {
	expected := createTestPDIntegration(t)
	defer cleanUpPDIntegration(t)

	actual, err := client.GetPDIntegration()
	if err != nil {
		t.Fatalf("Retrieving a pagerduty integration failed when it shouldn't: (%s)", err)
	}

	expectedServiceNames := make([]string, len(expected.Services))
	for _, service := range expected.Services {
		expectedServiceNames = append(expectedServiceNames, service.ServiceName)
	}

	actualServiceNames := make([]string, len(actual.Services))
	for _, service := range actual.Services {
		actualServiceNames = append(actualServiceNames, service.ServiceName)
	}

	assert.Equal(t, expectedServiceNames, actualServiceNames)
}

func TestIntegrationsPDUpdate(t *testing.T) {
	pdIntegration := createTestPDIntegration(t)
	defer cleanUpPDIntegration(t)

	pdIntegration.Services = append(pdIntegration.Services, datadog.PDService{
		ServiceName: "test-pd-update",
		ServiceKey:  "test-pd-update-key",
	})

	if err := client.UpdatePDIntegration(pdIntegration); err != nil {
		t.Fatalf("Updating a pagerduty integration failed when it shouldn't: %s", err)
	}

	actual, err := client.GetPDIntegration()
	if err != nil {
		t.Fatalf("Retrieving a pagerduty integration failed when it shouldn't: %s", err)
	}

	expectedServiceNames := make([]string, len(pdIntegration.Services))
	for _, service := range pdIntegration.Services {
		expectedServiceNames = append(expectedServiceNames, service.ServiceName)
	}

	actualServiceNames := make([]string, len(actual.Services))
	for _, service := range actual.Services {
		actualServiceNames = append(actualServiceNames, service.ServiceName)
	}

	assert.Equal(t, expectedServiceNames, actualServiceNames)
}

func TestIntegrationsPDGet(t *testing.T) {
	pdIntegration := createTestPDIntegration(t)
	defer cleanUpPDIntegration(t)

	actual, err := client.GetPDIntegration()
	if err != nil {
		t.Fatalf("Retrieving pdIntegration failed when it shouldn't: %s", err)
	}

	expectedServiceNames := make([]string, len(pdIntegration.Services))
	for _, service := range pdIntegration.Services {
		expectedServiceNames = append(expectedServiceNames, service.ServiceName)
	}

	actualServiceNames := make([]string, len(actual.Services))
	for _, service := range actual.Services {
		actualServiceNames = append(actualServiceNames, service.ServiceName)
	}

	assert.Equal(t, expectedServiceNames, actualServiceNames)
}

func getTestPDIntegration() *datadog.PDIntegration {
	return &datadog.PDIntegration{
		Services: []datadog.PDService{
			{
				ServiceName: "testPDServiceName",
				ServiceKey:  "testPDServiceKey",
			},
		},
		Subdomain: "testdomain",
		// Datadog will actually validate this value
		// so we're leaving it blank for tests
		Schedules: []string{},
		APIToken:  "abc123",
		RunCheck:  false,
	}
}

func createTestPDIntegration(t *testing.T) *datadog.PDIntegration {
	pdIntegration := getTestPDIntegration()
	err := client.CreatePDIntegration(pdIntegration)
	if err != nil {
		t.Fatalf("Creating a pdIntegration failed when it shouldn't: %s", err)
	}

	return pdIntegration
}

func cleanUpPDIntegration(t *testing.T) {
	if err := client.DeletePDIntegration(); err != nil {
		t.Fatalf("Deleting the pagerduty integration failed when it shouldn't. Manual cleanup needed. (%s)", err)
	}

	pdIntegration, err := client.GetPDIntegration()
	if pdIntegration != nil {
		t.Fatal("PagerDuty Integration hasn't been deleted when it should have been. Manual cleanup needed.")
	}

	if err == nil {
		t.Fatal("Fetching deleted pagerduty integration didn't lead to an error.")
	}
}
