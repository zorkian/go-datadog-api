package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
)

func TestDowntimeCreateAndDelete(t *testing.T) {
	expected := getTestDowntime()
	// create the downtime and compare it
	actual := createTestDowntime(t)
	defer cleanUpDowntime(t, *actual.Id)

	// Set ID of our original struct to ID of the returned struct so we can easily compare the results
	expected.SetId(actual.GetId())
	// Set creator ID for the same reason (there's no easy way to get ID of current user ATM,
	// but if there was, we could do this dynamically)
	expected.SetCreatorID(actual.GetCreatorID())

	assert.Equal(t, expected, actual)

	actual, err := client.GetDowntime(*actual.Id)
	if err != nil {
		t.Fatalf("Retrieving a downtime failed when it shouldn't: (%s)", err)
	}
	assert.Equal(t, expected, actual)
}

func TestDowntimeLinkedToMonitorCreateAndDelete(t *testing.T) {
	monitor := createTestMonitor(t)
	defer cleanUpMonitor(t, monitor.GetId())

	expected := getTestDowntime()
	expected.SetMonitorId(monitor.GetId())

	downtime, err := client.CreateDowntime(expected)
	defer cleanUpDowntime(t, downtime.GetId())
	if err != nil {
		t.Fatalf("Creating a downtime failed when it shouldn't: %s", err)
	}

	expected.SetId(downtime.GetId())
	expected.SetCreatorID(downtime.GetCreatorID())

	assert.Equal(t, expected, downtime)

	actual, err := client.GetDowntime(downtime.GetId())
	if err != nil {
		t.Fatalf("Retrieving a downtime failed when it shouldn't: (%s)", err)
	}
	assert.Equal(t, expected, actual)
}

func TestDowntimeUpdate(t *testing.T) {

	downtime := createTestDowntime(t)
	originalID := int(downtime.GetId())

	// changing the scope will cause the downtime to be replaced in the future
	// this test should be updated to validate this
	downtime.Scope = []string{"env:downtime_test", "env:downtime_test2"}
	defer cleanUpDowntime(t, *downtime.Id)

	if err := client.UpdateDowntime(downtime); err != nil {
		t.Fatalf("Updating a downtime failed when it shouldn't: %s", err)
	}

	actual, err := client.GetDowntime(*downtime.Id)
	if err != nil {
		t.Fatalf("Retrieving a downtime failed when it shouldn't: %s", err)
	}

	// uncomment once immutable to validate it changed to NotEqual
	assert.Equal(t, originalID, actual.GetId())
	assert.Equal(t, downtime.GetActive(), actual.GetActive())
	assert.Equal(t, downtime.GetCanceled(), actual.GetCanceled())
	assert.Equal(t, downtime.GetDisabled(), actual.GetDisabled())
	assert.Equal(t, downtime.GetEnd(), actual.GetEnd())
	assert.Equal(t, downtime.GetMessage(), actual.GetMessage())
	assert.Equal(t, downtime.GetMonitorId(), actual.GetMonitorId())
	assert.Equal(t, downtime.MonitorTags, actual.MonitorTags)
	assert.Equal(t, downtime.GetParentId(), actual.GetParentId())
	// timezone will be automatically updated to UTC
	assert.Equal(t, "UTC", actual.GetTimezone())
	assert.Equal(t, downtime.GetRecurrence(), actual.GetRecurrence())
	assert.EqualValues(t, downtime.Scope, actual.Scope)
	// in the future the replaced downtime will have an updated start time, flip this to NotEqual
	assert.Equal(t, downtime.GetStart(), actual.GetStart())
}

func TestDowntimeGet(t *testing.T) {
	downtimes, err := client.GetDowntimes()
	if err != nil {
		t.Fatalf("Retrieving downtimes failed when it shouldn't: %s", err)
	}
	num := len(downtimes)

	downtime := createTestDowntime(t)
	defer cleanUpDowntime(t, *downtime.Id)

	downtimes, err = client.GetDowntimes()
	if err != nil {
		t.Fatalf("Retrieving downtimes failed when it shouldn't: %s", err)
	}

	if num+1 != len(downtimes) {
		t.Fatalf("Number of downtimes didn't match expected: %d != %d", len(downtimes), num+1)
	}
}

func getTestDowntime() *datadog.Downtime {

	r := &datadog.Recurrence{
		Type:     datadog.String("weeks"),
		Period:   datadog.Int(1),
		WeekDays: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
	}

	return &datadog.Downtime{
		Active:      datadog.Bool(false),
		CreatorID:   datadog.Int(123),
		Disabled:    datadog.Bool(false),
		Message:     datadog.String("Test downtime message"),
		MonitorTags: []string{"some:tag"},
		ParentId:    nil,
		Timezone:    datadog.String("UTC"),
		Scope:       []string{"env:downtime_test"},
		Start:       datadog.Int(1577836800),
		End:         datadog.Int(1577840400),
		Recurrence:  r,
		Type:        datadog.Int(2),
	}
}

func createTestDowntime(t *testing.T) *datadog.Downtime {
	downtime := getTestDowntime()
	downtime, err := client.CreateDowntime(downtime)
	if err != nil {
		t.Fatalf("Creating a downtime failed when it shouldn't: %s", err)
	}

	return downtime
}

func cleanUpDowntime(t *testing.T, id int) {
	if err := client.DeleteDowntime(id); err != nil {
		t.Fatalf("Deleting a downtime failed when it shouldn't. Manual cleanup needed. (%s)", err)
	}

	deletedDowntime, err := client.GetDowntime(id)
	if deletedDowntime != nil && *deletedDowntime.Canceled == 0 {
		t.Fatal("Downtime hasn't been deleted when it should have been. Manual cleanup needed.")
	}

	if err == nil && *deletedDowntime.Canceled == 0 {
		t.Fatal("Fetching deleted downtime didn't lead to an error and downtime Canceled not set.")
	}
}
