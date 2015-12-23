// +build integration

package datadog

import (
	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
	"testing"
)

func init() {
	client = initTest()
}

func TestCreateAndDeleteMonitor(t *testing.T) {
	expected := getTestMonitor()
	// create the monitor and compare it
	actual := createTestMonitor(t)
	defer cleanUpMonitor(t, actual.Id)

	// the ID of our original struct will be zero. To prevent having to loop through
	// all other values, set the ID of expected to what we created
	expected.Id = actual.Id
	assert.Equal(t, actual, expected)

	// now try to fetch it freshly and compare it again
	actual, err := client.GetMonitor(actual.Id)
	if err != nil {
		t.Fatalf("Retrieving a monitor failed when it shouldn't. Manual needed (%s)", err)
	}
	assert.Equal(t, actual, expected)
}

func TestUpdateMonitor(t *testing.T) {

	monitor := createTestMonitor(t)
	defer cleanUpMonitor(t, monitor.Id)

	monitor.Name = "___New-Test-Monitor___"
	if err := client.UpdateMonitor(monitor); err != nil {
		t.Fatalf("Updating a monitor failed when it shouldn't: %s", err)
	}

	actual, err := client.GetMonitor(monitor.Id)
	if err != nil {
		t.Fatalf("Retreiving a monitor failed when it shouldn't: %s", err)
	}

	assert.Equal(t, actual, monitor)

}

func TestGetMonitor(t *testing.T) {
	monitors, err := client.GetMonitors()
	if err != nil {
		t.Fatalf("Retreiving monitors failed when it shouldn't: %s", err)
	}
	num := len(monitors)

	monitor := createTestMonitor(t)
	defer cleanUpMonitor(t, monitor.Id)

	monitors, err = client.GetMonitors()
	if err != nil {
		t.Fatalf("Retreiving monitors failed when it shouldn't: %s", err)
	}

	if num+1 != len(monitors) {
		t.Fatalf("Number of monitors didn't match expected: %d != %d", len(monitors), num+1)
	}
}

/*
	TODO: add
	MuteMonitors
	UnmuteMonitors
	MuteMonitor
	UnmuteMonitor
*/

func getTestMonitor() *datadog.Monitor {

	o := datadog.Options{
		NotifyNoData:    true,
		NoDataTimeframe: 60,
		Silenced:        map[string]int{}, // TODO can we make it so that the library inits this by default?
	}

	return &datadog.Monitor{
		Message: "Test message",
		Query:   "avg(last_15m):avg:system.disk.in_use{*} by {host,device} > 0.8",
		Name:    "Test monitor",
		Options: o,
		Type:    "metric alert",
	}
}

func createTestMonitor(t *testing.T) *datadog.Monitor {
	monitor := getTestMonitor()
	monitor, err := client.CreateMonitor(monitor)
	if err != nil {
		t.Fatalf("Creating a monitor failed when it shouldn't: %s", err)
	}

	return monitor
}

func cleanUpMonitor(t *testing.T, id int) {
	if err := client.DeleteMonitor(id); err != nil {
		t.Fatalf("Deleting a monitor failed when it shouldn't. Manual cleanup needed. (%s)", err)
	}

	deletedMonitor, err := client.GetMonitor(id)
	// TODO check if it's a 404, handle a non 404 a bit differently
	if deletedMonitor != nil {
		t.Fatal("Monitor hasn't been deleted when it should have been. Manual cleanup needed.")
	}

	if err == nil {
		t.Fatal("Fetching deleted monitor didn't lead to an error.")
	}
}
