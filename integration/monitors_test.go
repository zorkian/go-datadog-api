package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
)

func TestMonitorCreateAndDelete(t *testing.T) {
	expected := getTestMonitor()
	// create the monitor and compare it
	actual := createTestMonitor(t)
	defer cleanUpMonitor(t, actual.GetId())

	// Set ID of our original struct to zero so we can easily compare the results
	expected.SetId(actual.GetId())
	// Set Creator to the original struct as we can't predict details of the creator
	expected.SetCreator(actual.GetCreator())

	assert.Equal(t, expected, actual)

	actual, err := client.GetMonitor(*actual.Id)
	if err != nil {
		t.Fatalf("Retrieving a monitor failed when it shouldn't: (%s)", err)
	}
	assert.Equal(t, expected, actual)
}

func TestMonitorUpdate(t *testing.T) {

	monitor := createTestMonitor(t)
	defer cleanUpMonitor(t, *monitor.Id)

	monitor.SetName("___New-Test-Monitor___")
	if err := client.UpdateMonitor(monitor); err != nil {
		t.Fatalf("Updating a monitor failed when it shouldn't: %s", err)
	}

	actual, err := client.GetMonitor(*monitor.Id)
	if err != nil {
		t.Fatalf("Retrieving a monitor failed when it shouldn't: %s", err)
	}

	assert.Equal(t, monitor, actual)

}

func TestMonitorUpdateRemovingTags(t *testing.T) {

	monitor := createTestMonitorWithTags(t)
	defer cleanUpMonitor(t, *monitor.Id)

	monitor.Tags = make([]string, 0)
	if err := client.UpdateMonitor(monitor); err != nil {
		t.Fatalf("Updating a monitor failed when it shouldn't: %s", err)
	}

	actual, err := client.GetMonitor(*monitor.Id)
	if err != nil {
		t.Fatalf("Retrieving a monitor failed when it shouldn't: %s", err)
	}

	assert.Equal(t, monitor, actual)

}

func TestMonitorGet(t *testing.T) {
	monitors, err := client.GetMonitors()
	if err != nil {
		t.Fatalf("Retrieving monitors failed when it shouldn't: %s", err)
	}
	num := len(monitors)

	monitor := createTestMonitor(t)
	defer cleanUpMonitor(t, *monitor.Id)

	monitors, err = client.GetMonitors()
	if err != nil {
		t.Fatalf("Retrieving monitors failed when it shouldn't: %s", err)
	}

	if num+1 != len(monitors) {
		t.Fatalf("Number of monitors didn't match expected: %d != %d", len(monitors), num+1)
	}
}

func TestMonitorGetWithoutNoDataTimeframe(t *testing.T) {
	monitors, err := client.GetMonitors()
	if err != nil {
		t.Fatalf("Retrieving monitors failed when it shouldn't: %s", err)
	}
	num := len(monitors)

	monitor := createTestMonitorWithoutNoDataTimeframe(t)
	defer cleanUpMonitor(t, *monitor.Id)

	monitors, err = client.GetMonitors()
	if err != nil {
		t.Fatalf("Retrieving monitors failed when it shouldn't: %s", err)
	}

	if num+1 != len(monitors) {
		t.Fatalf("Number of monitors didn't match expected: %d != %d", len(monitors), num+1)
	}
}

func TestMonitorMuteUnmute(t *testing.T) {
	monitor := createTestMonitor(t)
	defer cleanUpMonitor(t, *monitor.Id)

	// Mute
	err := client.MuteMonitor(*monitor.Id)
	if err != nil {
		t.Fatalf("Failed to mute monitor")

	}

	monitor, err = client.GetMonitor(*monitor.Id)
	if err != nil {
		t.Fatalf("Retrieving monitors failed when it shouldn't: %s", err)
	}

	// Mute without options will result in monitor.Options.Silenced
	// to have a key of "*" with value 0
	assert.Equal(t, 0, monitor.Options.Silenced["*"])

	// Unmute
	err = client.UnmuteMonitor(*monitor.Id)
	if err != nil {
		t.Fatalf("Failed to unmute monitor")
	}

	// Update remote state
	monitor, err = client.GetMonitor(*monitor.Id)
	if err != nil {
		t.Fatalf("Retrieving monitors failed when it shouldn't: %s", err)
	}

	// Assert this map is empty
	assert.Equal(t, 0, len(monitor.Options.Silenced))
}

func TestMonitorMuteUnmuteScopes(t *testing.T) {
	monitor := createTestMonitor(t)
	err := client.MuteMonitorScope(*monitor.Id, &datadog.MuteMonitorScope{Scope: datadog.String("host:foo")})
	if err != nil {
		t.Fatalf("Failed to mute monitor scope host:foo: %s", err)
	}

	err = client.MuteMonitorScope(*monitor.Id, &datadog.MuteMonitorScope{Scope: datadog.String("host:bar")})
	if err != nil {
		t.Fatalf("Failed to mute monitor scope host:bar: %s", err)
	}

	monitor, err = client.GetMonitor(*monitor.Id)
	if err != nil {
		t.Fatalf("Retrieving monitor failed when it shouldn't: %s", err)
	}
	assert.Equal(t, monitor.Options.Silenced, map[string]int{"host:foo": 0, "host:bar": 0})

	err = client.UnmuteMonitorScopes(*monitor.Id, &datadog.UnmuteMonitorScopes{Scope: datadog.String("host:foo")})
	if err != nil {
		t.Fatalf("Failed to unmute monitor scope host:foo: %s", err)
	}

	monitor, err = client.GetMonitor(*monitor.Id)
	if err != nil {
		t.Fatalf("Retrieving monitor failed when it shouldn't: %s", err)
	}
	assert.Equal(t, monitor.Options.Silenced, map[string]int{"host:bar": 0})

	err = client.UnmuteMonitorScopes(*monitor.Id, &datadog.UnmuteMonitorScopes{AllScopes: datadog.Bool(true)})
	if err != nil {
		t.Fatalf("Failed to unmute all monitor scopes: %s", err)
	}

	monitor, err = client.GetMonitor(*monitor.Id)
	if err != nil {
		t.Fatalf("Retrieving monitor failed when it shouldn't: %s", err)
	}
	assert.Equal(t, monitor.Options.Silenced, map[string]int{})
}

/*
	Testing of global mute and unmuting has not been added for following reasons:
	* Disabling and enabling of global monitoring does an @all mention which is noisy
	* It exposes risk to users that run integration tests in their main account
	* There is no endpoint to verify success
*/

func getTestMonitor() *datadog.Monitor {

	o := &datadog.Options{
		NotifyNoData:      datadog.Bool(true),
		NotifyAudit:       datadog.Bool(false),
		Locked:            datadog.Bool(false),
		NoDataTimeframe:   60,
		NewHostDelay:      datadog.Int(600),
		RequireFullWindow: datadog.Bool(true),
		Silenced:          map[string]int{},
		IncludeTags:       datadog.Bool(false),
	}

	return &datadog.Monitor{
		Message:      datadog.String("Test message"),
		Query:        datadog.String("avg(last_15m):avg:system.disk.in_use{*} by {host,device} > 0.8"),
		Name:         datadog.String("Test monitor"),
		Options:      o,
		Type:         datadog.String("metric alert"),
		Tags:         make([]string, 0),
		OverallState: datadog.String("No Data"),
	}
}

func getTestMonitorWithTags() *datadog.Monitor {

	o := &datadog.Options{
		NotifyNoData:      datadog.Bool(true),
		NotifyAudit:       datadog.Bool(false),
		Locked:            datadog.Bool(false),
		NoDataTimeframe:   60,
		NewHostDelay:      datadog.Int(600),
		RequireFullWindow: datadog.Bool(true),
		Silenced:          map[string]int{},
		IncludeTags:       datadog.Bool(true),
	}

	return &datadog.Monitor{
		Message: datadog.String("Test message"),
		Query:   datadog.String("avg(last_15m):avg:system.disk.in_use{*} by {host,device} > 0.8"),
		Name:    datadog.String("Test monitor"),
		Options: o,
		Type:    datadog.String("metric alert"),
		Tags:    []string{"foo:bar", "bar:baz"},
	}
}

func getTestMonitorWithoutNoDataTimeframe() *datadog.Monitor {

	o := &datadog.Options{
		NotifyNoData:      datadog.Bool(false),
		NotifyAudit:       datadog.Bool(false),
		Locked:            datadog.Bool(false),
		NewHostDelay:      datadog.Int(600),
		RequireFullWindow: datadog.Bool(true),
		Silenced:          map[string]int{},
		IncludeTags:       datadog.Bool(false),
	}

	return &datadog.Monitor{
		Message: datadog.String("Test message"),
		Query:   datadog.String("avg(last_15m):avg:system.disk.in_use{*} by {host,device} > 0.8"),
		Name:    datadog.String("Test monitor"),
		Options: o,
		Type:    datadog.String("metric alert"),
		Tags:    make([]string, 0),
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

func createTestMonitorWithTags(t *testing.T) *datadog.Monitor {
	monitor := getTestMonitorWithTags()
	monitor, err := client.CreateMonitor(monitor)
	if err != nil {
		t.Fatalf("Creating a monitor failed when it shouldn't: %s", err)
	}

	return monitor
}

func createTestMonitorWithoutNoDataTimeframe(t *testing.T) *datadog.Monitor {
	monitor := getTestMonitorWithoutNoDataTimeframe()
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
	if deletedMonitor != nil {
		t.Fatal("Monitor hasn't been deleted when it should have been. Manual cleanup needed.")
	}

	if err == nil {
		t.Fatal("Fetching deleted monitor didn't lead to an error.")
	}
}
