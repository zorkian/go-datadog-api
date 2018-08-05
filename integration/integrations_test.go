package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
)

/*
	PagerDuty Integration
*/

func TestIntegrationPDCreateAndDelete(t *testing.T) {
	expected := createTestIntegrationPD(t)
	defer cleanUpIntegrationPD(t)

	actual, err := client.GetIntegrationPD()
	if err != nil {
		t.Fatalf("Retrieving a PagerDuty integration failed when it shouldn't: (%s)", err)
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
		t.Fatalf("Updating a PagerDuty integration failed when it shouldn't: %s", err)
	}

	actual, err := client.GetIntegrationPD()
	if err != nil {
		t.Fatalf("Retrieving a PagerDuty integration failed when it shouldn't: %s", err)
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
		Schedules: []string{},
		APIToken:  datadog.String("abc123"),
	}
}

func createTestIntegrationPD(t *testing.T) *datadog.IntegrationPDRequest {
	pdIntegration := getTestIntegrationPD()
	err := client.CreateIntegrationPD(pdIntegration)
	if err != nil {
		t.Fatalf("Creating a PagerDuty integration failed when it shouldn't: %s", err)
	}

	return pdIntegration
}

func cleanUpIntegrationPD(t *testing.T) {
	if err := client.DeleteIntegrationPD(); err != nil {
		t.Fatalf("Deleting the PagerDuty integration failed when it shouldn't. Manual cleanup needed. (%s)", err)
	}

	pdIntegration, err := client.GetIntegrationPD()
	if pdIntegration != nil {
		t.Fatal("PagerDuty Integration hasn't been deleted when it should have been. Manual cleanup needed.")
	}

	if err == nil {
		t.Fatal("Fetching deleted PagerDuty integration didn't lead to an error.")
	}
}

/*
	Slack Integration
*/

func TestIntegrationSlackCreateAndDelete(t *testing.T) {
	expected := createTestIntegrationSlack(t)
	defer cleanUpIntegrationSlack(t)

	actual, err := client.GetIntegrationSlack()
	if err != nil {
		t.Fatalf("Retrieving a Slack integration failed when it shouldn't: (%s)", err)
	}

	expectedServiceHooksAccounts := make([]*string, len(expected.ServiceHooks))
	for _, service := range expected.ServiceHooks {
		expectedServiceHooksAccounts = append(expectedServiceHooksAccounts, service.Account)
	}

	actualServiceHooksAccounts := make([]*string, len(actual.ServiceHooks))
	for _, service := range actual.ServiceHooks {
		actualServiceHooksAccounts = append(actualServiceHooksAccounts, service.Account)
	}

	assert.Equal(t, expectedServiceHooksAccounts, actualServiceHooksAccounts)
}

func TestIntegrationSlackUpdate(t *testing.T) {
	slackIntegration := createTestIntegrationSlack(t)
	defer cleanUpIntegrationSlack(t)

	slackIntegration.ServiceHooks = append(slackIntegration.ServiceHooks, datadog.ServiceHookSlackRequest{
		Account: datadog.String("Main_Account_2"),
		Url:     datadog.String("https://hooks.slack.com/services/2/2"),
	})

	if err := client.UpdateIntegrationSlack(slackIntegration); err != nil {
		t.Fatalf("Updating a Slack integration failed when it shouldn't: %s", err)
	}

	actual, err := client.GetIntegrationSlack()
	if err != nil {
		t.Fatalf("Retrieving a Slack integration failed when it shouldn't: %s", err)
	}

	expectedServiceHooksAccounts := make([]*string, len(slackIntegration.ServiceHooks))
	for _, service := range slackIntegration.ServiceHooks {
		expectedServiceHooksAccounts = append(expectedServiceHooksAccounts, service.Account)
	}

	actualServiceHooksAccounts := make([]*string, len(actual.ServiceHooks))
	for _, service := range actual.ServiceHooks {
		actualServiceHooksAccounts = append(actualServiceHooksAccounts, service.Account)
	}

	assert.Equal(t, expectedServiceHooksAccounts, actualServiceHooksAccounts)
}

func TestIntegrationSlackGet(t *testing.T) {
	slackIntegration := createTestIntegrationSlack(t)
	defer cleanUpIntegrationSlack(t)

	actual, err := client.GetIntegrationSlack()
	if err != nil {
		t.Fatalf("Retrieving Slack integration failed when it shouldn't: %s", err)
	}

	expectedServiceHooksAccounts := make([]*string, len(slackIntegration.ServiceHooks))
	for _, service := range slackIntegration.ServiceHooks {
		expectedServiceHooksAccounts = append(expectedServiceHooksAccounts, service.Account)
	}

	actualServiceHooksAccounts := make([]*string, len(actual.ServiceHooks))
	for _, service := range actual.ServiceHooks {
		actualServiceHooksAccounts = append(actualServiceHooksAccounts, service.Account)
	}

	assert.Equal(t, expectedServiceHooksAccounts, actualServiceHooksAccounts)
}

func getTestIntegrationSlack() *datadog.IntegrationSlackRequest {
	return &datadog.IntegrationSlackRequest{
		ServiceHooks: []datadog.ServiceHookSlackRequest{
			{
				Account: datadog.String("Main_Account"),
				Url:     datadog.String("https://hooks.slack.com/services/1/1"),
			},
		},
		Channels: []datadog.ChannelSlackRequest{
			{
				ChannelName:             datadog.String("#private"),
				TransferAllUserComments: datadog.Bool(true),
				Account:                 datadog.String("Main_Account"),
			},
		},
	}
}

func createTestIntegrationSlack(t *testing.T) *datadog.IntegrationSlackRequest {
	slackIntegration := getTestIntegrationSlack()

	err := client.CreateIntegrationSlack(slackIntegration)
	if err != nil {
		t.Fatalf("Creating a Slack integration failed when it shouldn't: %s", err)
	}

	return slackIntegration
}

func cleanUpIntegrationSlack(t *testing.T) {
	if err := client.DeleteIntegrationSlack(); err != nil {
		t.Fatalf("Deleting the Slack integration failed when it shouldn't. Manual cleanup needed. (%s)", err)
	}

	slackIntegration, err := client.GetIntegrationSlack()
	if slackIntegration != nil {
		t.Fatal("Slack Integration hasn't been deleted when it should have been. Manual cleanup needed.")
	}

	if err == nil {
		t.Fatal("Fetching deleted Slack integration didn't lead to an error.")
	}
}

/*
	AWS Integration
*/

func TestIntegrationAWSCreateAndDelete(t *testing.T) {
	// Validate creation of AWS Integration
	expected := getTestIntegrationAWS()
	createAwsIntegrationResponse := createTestIntegrationAWS(t)
	defer cleanUpIntegrationAWS(t, getTestIntegrationAWSDeleteRequest())

	assert.NotNil(t, createAwsIntegrationResponse.ExternalID, "An external ID should have been returned from Datadog on integration create.")

	// Get the AWS Accounts from the AWS Integration
	awsAccountsInDatadog, err := client.GetIntegrationAWS()
	if err != nil {
		t.Fatalf("Retrieving a AWS integration failed when it shouldn't: (%s)", err)
	}

	// Check the created AWS Account is in the slice
	var createdAWSAccount datadog.IntegrationAWSAccount
	for _, account := range *awsAccountsInDatadog {
		if *account.AccountID == *expected.AccountID {
			createdAWSAccount = account
		}
	}

	// Test each property as slices order can change in FilterTags and HostTags
	assert.Equal(t, expected.AccountID, createdAWSAccount.AccountID)
	assert.Equal(t, expected.RoleName, createdAWSAccount.RoleName)
	assert.ElementsMatch(t, expected.FilterTags, createdAWSAccount.FilterTags)
	assert.ElementsMatch(t, expected.HostTags, createdAWSAccount.HostTags)
	assert.Equal(t, expected.AccountSpecificNamespaceRules, createdAWSAccount.AccountSpecificNamespaceRules)
}

func getTestIntegrationAWS() *datadog.IntegrationAWSAccount {
	return &datadog.IntegrationAWSAccount{
		AccountID: datadog.String("1111111111111"),
		RoleName:  datadog.String("GoLangDatadogAWSIntegrationRole"),
		FilterTags: []string{
			"env:production",
			"instance-type:c1.*",
			"!region:us-east-1",
		},
		HostTags: []string{
			"account:my_aws_account",
		},
		AccountSpecificNamespaceRules: map[string]bool{
			"auto_scaling": false,
			"opsworks":     false,
		},
	}
}

func getTestIntegrationAWSDeleteRequest() *datadog.IntegrationAWSAccountDeleteRequest {
	return &datadog.IntegrationAWSAccountDeleteRequest{
		AccountID: datadog.String("1111111111111"),
		RoleName:  datadog.String("GoLangDatadogAWSIntegrationRole"),
	}
}

func createTestIntegrationAWS(t *testing.T) *datadog.IntegrationAWSAccountCreateResponse {
	awsIntegrationRequest := getTestIntegrationAWS()

	result, err := client.CreateIntegrationAWS(awsIntegrationRequest)
	if err != nil {
		t.Fatalf("Creating a AWS integration failed when it shouldn't: %s", err.Error())
	}

	return result
}

func cleanUpIntegrationAWS(t *testing.T, awsAccount *datadog.IntegrationAWSAccountDeleteRequest) {
	if err := client.DeleteIntegrationAWS(awsAccount); err != nil {
		t.Fatalf("Deleting the AWS Account from the AWS integration failed when it shouldn't. Manual cleanup needed. (%s)", err)
	}

	awsIntegration, err := client.GetIntegrationAWS()

	// check the account is no longer in Datadog
	for _, account := range *awsIntegration {
		if *account.AccountID == *awsAccount.AccountID {
			t.Fatal("AWS Account in the AWS Integration hasn't been deleted when it should have been. Manual cleanup needed.")
		}
	}

	if err != nil {
		t.Fatalf("Getting AWS accounts from the AWS integration failed when it shouldn't: %s", err)
	}
}
