package integration

import (
	"testing"

	datadog "github.com/zorkian/go-datadog-api"
)

func TestDashboardListItemsV2GetAndUpdate(t *testing.T) {
	list := getTestDashboardList()
	list, err := client.CreateDashboardList(list)
	if err != nil {
		t.Fatalf("Creating a dashboard list failed when it shouldn't: %s", err)
	}

	defer cleanUpDashboardList(t, *list.Id)

	if *list.DashboardCount != 0 {
		t.Fatalf("Number of dashboards in dashboard list does not match: %d != %d", *list.DashboardCount, 0)
	}

	// Test update with timeboard in list
	timeboard := createTestDashboard(t)
	defer cleanUpDashboard(t, *timeboard.Id)

	timeboardItems := []datadog.DashboardListV2Item{
		getTestDashboardListItemV2Timeboard(*timeboard.NewId),
	}

	actualItems, err := client.UpdateDashboardListV2Items(*list.Id, timeboardItems)
	if err != nil {
		t.Fatalf("Updating dashboard list items failed when it shouldn't: %s", err)
	}
	if len(actualItems) != len(timeboardItems) {
		t.Fatalf("Number of updated dashboards does not match: %d != %d", len(actualItems), len(timeboardItems))
	}
	assertDashboardListItemV2Equals(t, &actualItems[0], &timeboardItems[0])

	// Get dashboard list to make sure count is 1
	list, err = client.GetDashboardList(*list.Id)
	if err != nil {
		t.Fatalf("Getting a dashboard list failed when it shouldn't: %s", err)
	}

	if *list.DashboardCount != len(timeboardItems) {
		t.Fatalf("Number of dashboards in dashboard list does not match: %d != %d", *list.DashboardCount, 0)
	}

	// Test update with screenboard in list
	screenboard := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *screenboard.Id)

	screenboardItems := []datadog.DashboardListV2Item{
		getTestDashboardListItemV2Screenboard(*screenboard.NewId),
	}

	actualItems, err = client.UpdateDashboardListV2Items(*list.Id, screenboardItems)
	if err != nil {
		t.Fatalf("Updating dashboard list items failed when it shouldn't: %s", err)
	}
	if len(actualItems) != len(screenboardItems) {
		t.Fatalf("Number of updated dashboards does not match: %d != %d", len(actualItems), len(screenboardItems))
	}
	assertDashboardListItemV2Equals(t, &actualItems[0], &screenboardItems[0])

	// Get dashboard list to make sure count is 1
	list, err = client.GetDashboardList(*list.Id)
	if err != nil {
		t.Fatalf("Getting a dashboard list failed when it shouldn't: %s", err)
	}

	if *list.DashboardCount != len(screenboardItems) {
		t.Fatalf("Number of dashboards in dashboard list does not match: %d != %d", *list.DashboardCount, 0)
	}
}

func TestDashboardListItemsV2AddAndDelete(t *testing.T) {
	list := getTestDashboardList()
	list, err := client.CreateDashboardList(list)
	if err != nil {
		t.Fatalf("Creating a dashboard list failed when it shouldn't: %s", err)
	}

	defer cleanUpDashboardList(t, *list.Id)

	// Add timeboard to list
	timeboard := createTestDashboard(t)
	defer cleanUpDashboard(t, *timeboard.Id)

	items := []datadog.DashboardListV2Item{
		getTestDashboardListItemV2Timeboard(*timeboard.NewId),
	}

	addedItems, err := client.AddDashboardListV2Items(*list.Id, items)
	if err != nil {
		t.Fatalf("Adding dashboard list items failed when it shouldn't: %s", err)
	}
	if len(addedItems) != len(items) {
		t.Fatalf("Number of updated dashboards does not match: %d != %d", len(addedItems), len(items))
	}
	assertDashboardListItemV2Equals(t, &addedItems[0], &items[0])

	// Get dashboard list to make sure count is 1
	list, err = client.GetDashboardList(*list.Id)
	if err != nil {
		t.Fatalf("Getting a dashboard list failed when it shouldn't: %s", err)
	}

	if *list.DashboardCount != len(items) {
		t.Fatalf("Number of dashboards in dashboard list does not match: %d != %d", *list.DashboardCount, 0)
	}

	// Adding an existing item should be ignored, meaning we should only get
	// one item back in the list of added items.
	screenboard := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *screenboard.Id)

	items = append(items, getTestDashboardListItemV2Screenboard(*screenboard.NewId))

	addedItems, err = client.AddDashboardListV2Items(*list.Id, items)
	if err != nil {
		t.Fatalf("Adding dashboard list items failed when it shouldn't: %s", err)
	}
	if len(addedItems) != 1 {
		t.Fatalf("Number of updated dashboards does not match: %d != %d", len(addedItems), 1)
	}
	assertDashboardListItemV2Equals(t, &addedItems[0], &items[1])

	// Get dashboard list to make sure count is now 2
	list, err = client.GetDashboardList(*list.Id)
	if err != nil {
		t.Fatalf("Getting a dashboard list failed when it shouldn't: %s", err)
	}

	if *list.DashboardCount != len(items) {
		t.Fatalf("Number of dashboards in dashboard list does not match: %d != %d", *list.DashboardCount, 0)
	}

	// Delete everything in the dashboard list and check length of deleted items is 2
	deletedItems, err := client.DeleteDashboardListV2Items(*list.Id, items)
	if err != nil {
		t.Fatalf("Deleting dashboard list items failed when it shouldn't: %s", err)
	}
	if len(deletedItems) != len(items) {
		t.Fatalf("Number of deleted dashboards does not match: %d != %d", len(deletedItems), len(items))
	}

	// Check that the dashboard list is empty again
	list, err = client.GetDashboardList(*list.Id)
	if err != nil {
		t.Fatalf("Getting a dashboard list failed when it shouldn't: %s", err)
	}

	if *list.DashboardCount != 0 {
		t.Fatalf("Number of dashboards in dashboard list does not match: %d != %d", *list.DashboardCount, 0)
	}
}

func getTestDashboardListItemV2Timeboard(id string) datadog.DashboardListV2Item {
	return datadog.DashboardListV2Item{
		ID:   datadog.String(id),
		Type: datadog.String(datadog.DashboardListItemCustomTimeboard),
	}
}

func assertDashboardListItemV2Equals(t *testing.T, actual, expected *datadog.DashboardListV2Item) {
	if *actual.ID != *expected.ID {
		t.Errorf("Dashboard list item id does not match: %s != %s", *actual.ID, *expected.ID)
	}
	if *actual.Type != *expected.Type {
		t.Errorf("Dashboard list item type does not match: %s != %s", *actual.Type, *expected.Type)
	}
}

func getTestDashboardListItemV2Screenboard(id string) datadog.DashboardListV2Item {
	return datadog.DashboardListV2Item{
		ID:   datadog.String(id),
		Type: datadog.String(datadog.DashboardListItemCustomScreenboard),
	}
}
