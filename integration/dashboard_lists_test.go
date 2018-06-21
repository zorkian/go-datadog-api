package integration

import (
	"log"
	"testing"

	"github.com/zorkian/go-datadog-api"
)

func TestDashboardListsGet(t *testing.T) {
	lists, err := client.GetDashboardLists()
	if err != nil {
		log.Fatalf("fatal: %s\n", err)
	}

	num := len(lists)
	list := createTestDashboardList(t)
	defer cleanUpDashboardList(t, *list.Id)

	lists, err = client.GetDashboardLists()
	if err != nil {
		t.Fatalf("Retrieving dashboard lists failed when it shouldn't: %s", err)
	}

	if num+1 != len(lists) {
		t.Fatalf("Number of dashboard lists didn't match expected: %d != %d", len(lists), num+1)
	}
}

func TestDashboardListCreateAndDelete(t *testing.T) {
	expected := getTestDashboardList()
	// create the dashboard list and compare it
	actual, err := client.CreateDashboardList(expected)
	if err != nil {
		t.Fatalf("Creating a dashboard list failed when it shouldn't: %s", err)
	}

	defer cleanUpDashboardList(t, *actual.Id)

	assertDashboardListEquals(t, actual, expected)

	// now try to fetch it freshly and compare it again
	actual, err = client.GetDashboardList(*actual.Id)
	if err != nil {
		t.Fatalf("Getting a dashboard list failed when it shouldn't: %s", err)
	}
	assertDashboardListEquals(t, actual, expected)
}

func TestDashboardListUpdate(t *testing.T) {
	list := getTestDashboardList()
	expected, err := client.CreateDashboardList(list)
	if err != nil {
		t.Fatalf("Creating a dashboard list failed when it shouldn't: %s", err)
	}

	defer cleanUpDashboardList(t, *expected.Id)
	expected.Name = datadog.String("___New-Test-Dashboard-List___")

	if err := client.UpdateDashboardList(expected); err != nil {
		t.Fatalf("Updating a dashboard list failed when it shouldn't: %s", err)
	}

	actual, err := client.GetDashboardList(*expected.Id)
	if err != nil {
		t.Fatalf("Getting a dashboard list failed when it shouldn't: %s", err)
	}

	assertDashboardListEquals(t, actual, expected)
}

func TestDashboardListItemsGetAndUpdate(t *testing.T) {
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

	timeboardItems := []datadog.DashboardListItem{
		getTestDashboardListItemTimeboard(*timeboard.Id),
	}

	actualItems, err := client.UpdateDashboardListItems(*list.Id, timeboardItems)
	if err != nil {
		t.Fatalf("Updating dashboard list items failed when it shouldn't: %s", err)
	}
	if len(actualItems) != len(timeboardItems) {
		t.Fatalf("Number of updated dashboards does not match: %d != %d", len(actualItems), len(timeboardItems))
	}
	assertDashboardListItemEquals(t, &actualItems[0], &timeboardItems[0])

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

	screenboardItems := []datadog.DashboardListItem{
		getTestDashboardListItemScreenboard(*screenboard.Id),
	}

	actualItems, err = client.UpdateDashboardListItems(*list.Id, screenboardItems)
	if err != nil {
		t.Fatalf("Updating dashboard list items failed when it shouldn't: %s", err)
	}
	if len(actualItems) != len(screenboardItems) {
		t.Fatalf("Number of updated dashboards does not match: %d != %d", len(actualItems), len(screenboardItems))
	}
	assertDashboardListItemEquals(t, &actualItems[0], &screenboardItems[0])

	// Get dashboard list to make sure count is 1
	list, err = client.GetDashboardList(*list.Id)
	if err != nil {
		t.Fatalf("Getting a dashboard list failed when it shouldn't: %s", err)
	}

	if *list.DashboardCount != len(screenboardItems) {
		t.Fatalf("Number of dashboards in dashboard list does not match: %d != %d", *list.DashboardCount, 0)
	}
}

func TestDashboardListItemsAddAndDelete(t *testing.T) {
	list := getTestDashboardList()
	list, err := client.CreateDashboardList(list)
	if err != nil {
		t.Fatalf("Creating a dashboard list failed when it shouldn't: %s", err)
	}

	defer cleanUpDashboardList(t, *list.Id)

	// Add timeboard to list
	timeboard := createTestDashboard(t)
	defer cleanUpDashboard(t, *timeboard.Id)

	items := []datadog.DashboardListItem{
		getTestDashboardListItemTimeboard(*timeboard.Id),
	}

	addedItems, err := client.AddDashboardListItems(*list.Id, items)
	if err != nil {
		t.Fatalf("Adding dashboard list items failed when it shouldn't: %s", err)
	}
	if len(addedItems) != len(items) {
		t.Fatalf("Number of updated dashboards does not match: %d != %d", len(addedItems), len(items))
	}
	assertDashboardListItemEquals(t, &addedItems[0], &items[0])

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

	items = append(items, getTestDashboardListItemScreenboard(*screenboard.Id))

	addedItems, err = client.AddDashboardListItems(*list.Id, items)
	if err != nil {
		t.Fatalf("Adding dashboard list items failed when it shouldn't: %s", err)
	}
	if len(addedItems) != 1 {
		t.Fatalf("Number of updated dashboards does not match: %d != %d", len(addedItems), 1)
	}
	assertDashboardListItemEquals(t, &addedItems[0], &items[1])

	// Get dashboard list to make sure count is now 2
	list, err = client.GetDashboardList(*list.Id)
	if err != nil {
		t.Fatalf("Getting a dashboard list failed when it shouldn't: %s", err)
	}

	if *list.DashboardCount != len(items) {
		t.Fatalf("Number of dashboards in dashboard list does not match: %d != %d", *list.DashboardCount, 0)
	}

	// Delete everything in the dashboard list and check length of deleted items is 2
	deletedItems, err := client.DeleteDashboardListItems(*list.Id, items)
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

func getTestDashboardList() *datadog.DashboardList {
	return &datadog.DashboardList{
		Name: datadog.String("___Test-Dashboard-List___"),
	}
}

func getTestDashboardListItemTimeboard(id int) datadog.DashboardListItem {
	return datadog.DashboardListItem{
		Id:   datadog.Int(id),
		Type: datadog.String(datadog.DashboardListItemCustomTimeboard),
	}
}

func getTestDashboardListItemScreenboard(id int) datadog.DashboardListItem {
	return datadog.DashboardListItem{
		Id:   datadog.Int(id),
		Type: datadog.String(datadog.DashboardListItemCustomScreenboard),
	}
}

func createTestDashboardList(t *testing.T) *datadog.DashboardList {
	list := getTestDashboardList()
	list, err := client.CreateDashboardList(list)
	if err != nil {
		t.Fatalf("Creating a dashboard list failed when it shouldn't: %s", err)
	}

	return list
}

func cleanUpDashboardList(t *testing.T, id int) {
	if err := client.DeleteDashboardList(id); err != nil {
		t.Fatalf("Deleting a dashboard list failed when it shouldn't. Manual cleanup needed. (%s)", err)
	}

	deletedBoard, err := client.GetDashboardList(id)
	if deletedBoard != nil {
		t.Fatalf("Dashboard list hasn't been deleted when it should have been. Manual cleanup needed.")
	}

	if err == nil {
		t.Fatal("Fetching deleted dashboard list didn't lead to an error. Manual cleanup needed.")
	}
}

func assertDashboardListEquals(t *testing.T, actual, expected *datadog.DashboardList) {
	if *actual.Name != *expected.Name {
		t.Errorf("Dashboard list name does not match: %s != %s", *actual.Name, *expected.Name)
	}
	if actual.GetDashboardCount() != expected.GetDashboardCount() {
		t.Errorf("Dashboard list dashboard count does not match: %d != %d", *actual.DashboardCount, *expected.DashboardCount)
	}
}

func assertDashboardListItemEquals(t *testing.T, actual, expected *datadog.DashboardListItem) {
	if *actual.Id != *expected.Id {
		t.Errorf("Dashboard list item id does not match: %d != %d", *actual.Id, *expected.Id)
	}
	if *actual.Type != *expected.Type {
		t.Errorf("Dashboard list item type does not match: %s != %s", *actual.Type, *expected.Type)
	}
}
