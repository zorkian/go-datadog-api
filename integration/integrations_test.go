package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
)

func init() {
	client = initTest()
}

func TestIntegrationPDCreateAndDelete(t *testing.T) {
	expected := createTestIntegrationPD(t)
	defer cleanUpIntegrationPD(t)

	actual, err := client.GetIntegrationPD()
	if err != nil {
		t.Fatalf("Retrieving a pagerduty integration failed when it shouldn't: (%s)", err)
	}

	expectedServiceNames := make([]*string, len(expected.Services))
	for _, service := range expected.Services {
		expectedServiceNames = append(expectedServiceNames, service.ServiceName)
	}

	actualServiceNames := make([]*string, len(actual.Services))
	for _, service := range actual.Services {
		actualServiceNames = append(actualServiceNames, service.ServiceName)
	}

	assert.Equal(t, expectedServiceNames, actualServiceNames)
}

func TestIntegrationPDUpdate(t *testing.T) {
	pdIntegration := createTestIntegrationPD(t)
	defer cleanUpIntegrationPD(t)

	pdIntegration.Services = append(pdIntegration.Services, datadog.ServicePDRequest{
		ServiceName: datadog.String("test-pd-update"),
		ServiceKey:  datadog.String("test-pd-update-key"),
	})

	if err := client.UpdateIntegrationPD(pdIntegration); err != nil {
		t.Fatalf("Updating a pagerduty integration failed when it shouldn't: %s", err)
	}

	actual, err := client.GetIntegrationPD()
	if err != nil {
		t.Fatalf("Retrieving a pagerduty integration failed when it shouldn't: %s", err)
	}

	expectedServiceNames := make([]*string, len(pdIntegration.Services))
	for _, service := range pdIntegration.Services {
		expectedServiceNames = append(expectedServiceNames, service.ServiceName)
	}

	actualServiceNames := make([]*string, len(actual.Services))
	for _, service := range actual.Services {
		actualServiceNames = append(actualServiceNames, service.ServiceName)
	}

	assert.Equal(t, expectedServiceNames, actualServiceNames)
}

func TestIntegrationPDGet(t *testing.T) {
	pdIntegration := createTestIntegrationPD(t)
	defer cleanUpIntegrationPD(t)

	actual, err := client.GetIntegrationPD()
	if err != nil {
		t.Fatalf("Retrieving pdIntegration failed when it shouldn't: %s", err)
	}

	expectedServiceNames := make([]*string, len(pdIntegration.Services))
	for _, service := range pdIntegration.Services {
		expectedServiceNames = append(expectedServiceNames, service.ServiceName)
	}

	actualServiceNames := make([]*string, len(actual.Services))
	for _, service := range actual.Services {
		actualServiceNames = append(actualServiceNames, service.ServiceName)
	}

	assert.Equal(t, expectedServiceNames, actualServiceNames)
}

func getTestIntegrationPD() *datadog.IntegrationPDRequest {
	return &datadog.IntegrationPDRequest{
		Services: []datadog.ServicePDRequest{
			{
				ServiceName: datadog.String("testPDServiceName"),
				ServiceKey:  datadog.String("testPDServiceKey"),
			},
		},
		Subdomain: datadog.String("testdomain"),
		// Datadog will actually validate this value
		// so we're leaving it blank for tests
		Schedules: []*string{},
		APIToken:  datadog.String("abc123"),
	}
}

func createTestIntegrationPD(t *testing.T) *datadog.IntegrationPDRequest {
	pdIntegration := getTestIntegrationPD()
	err := client.CreateIntegrationPD(pdIntegration)
	if err != nil {
		t.Fatalf("Creating a pagerduty integration failed when it shouldn't: %s", err)
	}

	return pdIntegration
}

func cleanUpIntegrationPD(t *testing.T) {
	if err := client.DeleteIntegrationPD(); err != nil {
		t.Fatalf("Deleting the pagerduty integration failed when it shouldn't. Manual cleanup needed. (%s)", err)
	}

	pdIntegration, err := client.GetIntegrationPD()
	if pdIntegration != nil {
		t.Fatal("PagerDuty Integration hasn't been deleted when it should have been. Manual cleanup needed.")
	}

	if err == nil {
		t.Fatal("Fetching deleted pagerduty integration didn't lead to an error.")
	}
}
