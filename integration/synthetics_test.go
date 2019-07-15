package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	datadog "github.com/zorkian/go-datadog-api"
)

func TestSyntheticsCreateAndDelete(t *testing.T) {
	expected := getTestSynthetics()
	// create the monitor and compare it
	actual := createTestSynthetics(t)
	defer cleanUpSynthetics(t, actual.GetPublicId())

	// Set ID of our original struct to zero so we can easily compare the results
	expected.SetPublicId(actual.GetPublicId())
	// Set Creator to the original struct as we can't predict details of the creator
	expected.SetCreatedAt(actual.GetCreatedAt())
	expected.SetModifiedAt(actual.GetModifiedAt())
	expected.SetMonitorId(actual.GetMonitorId())

	assert.Equal(t, expected, actual)

	actual, err := client.GetSyntheticsTest(*actual.PublicId)
	if err != nil {
		t.Fatalf("Retrieving a synthetics failed when it shouldn't: (%s)", err)
	}
	expected.SetStatus(actual.GetStatus())
	expected.SetCreatedBy(actual.GetCreatedBy())
	expected.SetModifiedBy(actual.GetModifiedBy())
	expected.SetMonitorId(actual.GetMonitorId())
	assert.Equal(t, expected, actual)
}

func TestSyntheticsUpdate(t *testing.T) {

	syntheticsTest := createTestSynthetics(t)
	defer cleanUpSynthetics(t, syntheticsTest.GetPublicId())

	syntheticsTest.SetName("___New-Test-Synthetics___")
	publicId := syntheticsTest.GetPublicId()
	createdAt := syntheticsTest.GetCreatedAt()
	syntheticsTest.PublicId = nil
	syntheticsTest.CreatedAt = nil
	syntheticsTest.ModifiedAt = nil
	syntheticsTest.MonitorId = nil
	actual, err := client.UpdateSyntheticsTest(publicId, syntheticsTest)
	if err != nil {
		t.Fatalf("Updating a synthetics test failed when it shouldn't: %s", err)
	}

	syntheticsTest.SetPublicId(publicId)
	syntheticsTest.SetCreatedAt(createdAt)
	syntheticsTest.SetModifiedAt(actual.GetModifiedAt())
	syntheticsTest.SetMonitorId(actual.GetMonitorId())
	assert.Equal(t, syntheticsTest, actual)

}

func TestSyntheticsUpdateRemovingTags(t *testing.T) {

	syntheticsTest := createTestSynthetics(t)
	defer cleanUpSynthetics(t, syntheticsTest.GetPublicId())

	publicId := syntheticsTest.GetPublicId()
	createdAt := syntheticsTest.GetCreatedAt()
	syntheticsTest.PublicId = nil
	syntheticsTest.CreatedAt = nil
	syntheticsTest.ModifiedAt = nil
	syntheticsTest.MonitorId = nil
	syntheticsTest.Tags = []string{}
	actual, err := client.UpdateSyntheticsTest(publicId, syntheticsTest)
	if err != nil {
		t.Fatalf("Updating a synthetics test failed when it shouldn't: %s", err)
	}

	syntheticsTest.SetPublicId(publicId)
	syntheticsTest.SetCreatedAt(createdAt)
	syntheticsTest.SetModifiedAt(actual.GetModifiedAt())
	syntheticsTest.SetMonitorId(actual.GetMonitorId())
	assert.Equal(t, syntheticsTest, actual)

}

func TestSyntheticsGetAllTests(t *testing.T) {
	syntheticsTests, err := client.GetSyntheticsTestsByType("api")
	if err != nil {
		t.Fatalf("Retrieving synthetics tests failed when it shouldn't: %s", err)
	}
	num := len(syntheticsTests)

	syntheticsTest := createTestSynthetics(t)
	defer cleanUpSynthetics(t, syntheticsTest.GetPublicId())

	syntheticsTests, err = client.GetSyntheticsTestsByType("api")
	if err != nil {
		t.Fatalf("Retrieving synthetics tests failed when it shouldn't: %s", err)
	}

	if num+1 != len(syntheticsTests) {
		t.Fatalf("Number of synthetics tests didn't match expected: %d != %d", len(syntheticsTests), num+1)
	}
}

func TestMonitorPauseResume(t *testing.T) {
	syntheticsTest := createTestSynthetics(t)
	defer cleanUpSynthetics(t, syntheticsTest.GetPublicId())

	publicId := syntheticsTest.GetPublicId()

	// Pause SyntheticsTest
	_, err := client.PauseSyntheticsTest(publicId)
	if err != nil {
		t.Fatalf("Failed to pause test")
	}

	syntheticsTest, err = client.GetSyntheticsTest(publicId)
	if err != nil {
		t.Fatalf("Retrieving synthetics test failed when it shouldn't: %s", err)
	}

	assert.Equal(t, "paused", *syntheticsTest.Status)

	// Resume SyntheticsTest
	_, err = client.ResumeSyntheticsTest(publicId)
	if err != nil {
		t.Fatalf("Failed to resume synthetics test")
	}

	syntheticsTest, err = client.GetSyntheticsTest(publicId)
	if err != nil {
		t.Fatalf("Retrieving synthetics test failed when it shouldn't: %s", err)
	}

	assert.Equal(t, "live", *syntheticsTest.Status)
}

func TestSyntheticsGetAllLocations(t *testing.T) {
	syntheticsLocations, err := client.GetSyntheticsLocations()
	if err != nil {
		t.Fatalf("Retrieving synthetics locations failed when it shouldn't: %s", err)
	}
	num := len(syntheticsLocations)

	if num == 0 {
		t.Fatalf("Number of synthetics locations should be more than 0")
	}
}

func TestSyntheticsGetAllDevices(t *testing.T) {
	syntheticsDevices, err := client.GetSyntheticsBrowserDevices()
	if err != nil {
		t.Fatalf("Retrieving synthetics browser devices failed when it shouldn't: %s", err)
	}
	num := len(syntheticsDevices)

	if num == 0 {
		t.Fatalf("Number of synthetics devices should be more than 0")
	}
}

/*
	Testing of global mute and unmuting has not been added for following reasons:
	* Disabling and enabling of global monitoring does an @all mention which is noisy
	* It exposes risk to users that run integration tests in their main account
	* There is no endpoint to verify success
*/

func getTestSynthetics() *datadog.SyntheticsTest {
	c := &datadog.SyntheticsConfig{
		Request: &datadog.SyntheticsRequest{
			Method:  datadog.String("GET"),
			Url:     datadog.String("https://example.org"),
			Timeout: datadog.Int(30),
		},
		Assertions: []datadog.SyntheticsAssertion{{
			Type:     datadog.String("statusCode"),
			Operator: datadog.String("is"),
			Target:   float64(200),
		}},
	}
	o := &datadog.SyntheticsOptions{
		TickEvery: datadog.Int(60),
	}

	return &datadog.SyntheticsTest{
		Message:   datadog.String("Test message"),
		Name:      datadog.String("Test synthetics"),
		Config:    c,
		Options:   o,
		Locations: []string{"aws:eu-central-1"},
		Status:    datadog.String("live"),
		Type:      datadog.String("api"),
		Tags:      []string{"tag1:value1", "tag2:value2"},
	}
}

func createTestSynthetics(t *testing.T) *datadog.SyntheticsTest {
	synthetics := getTestSynthetics()
	synthetics, err := client.CreateSyntheticsTest(synthetics)
	if err != nil {
		t.Fatalf("Creating a synthetics failed when it shouldn't: %s", err)
	}

	return synthetics
}

func cleanUpSynthetics(t *testing.T, publicId string) {
	if err := client.DeleteSyntheticsTests([]string{publicId}); err != nil {
		t.Fatalf("Deleting a synthetics failed when it shouldn't. Manual cleanup needed. (%s)", err)
	}

	deletedSynthetics, err := client.GetSyntheticsTest(publicId)
	if deletedSynthetics != nil {
		t.Fatal("Synthetics hasn't been deleted when it should have been. Manual cleanup needed.")
	}

	if err == nil {
		t.Fatal("Fetching deleted synthetics didn't lead to an error.")
	}
}
