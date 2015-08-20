package integration_test

import (
	"log"
	"os"
	"testing"

	"github.com/ojongerius/go-datadog-api"
)

var (
	apiKey string
	appKey string
	client *datadog.Client
)

func init() {
	apiKey = os.Getenv("DATADOG_API_KEY")
	appKey = os.Getenv("DATADOG_APP_KEY")

	if apiKey == "" || appKey == "" {
		log.Fatal("Please make sure to set the env variables 'DATADOG_API_KEY' and 'DATADOG_APP_KEY' before running this test")
	}

	client = datadog.NewClient(apiKey, appKey)
}

func TestMain(m *testing.M) {
	num := countDashboards()

	result := m.Run()

	if num != countDashboards() {
		log.Fatal("Tests didn't clean-up all created dashboards. Manual clean-up required.")
	}

	os.Exit(result)
}

func TestCreateAndDeleteDashboard(t *testing.T) {
	expected := getTestDashboard()
	// create the dashboard and compare it
	actual, err := client.CreateDashboard(expected)
	if err != nil {
		t.Fatalf("Creating a dashboard failed when it shouldn't. (%s)", err)
	}
	assertDashboardEquals(t, actual, expected)

	// now try to fetch it freshly and compare it again
	actual, err = client.GetDashboard(actual.Id)
	if err != nil {
		t.Fatalf("Retreiving a dashboard failed when it shouldn't. (%s)", err)
	}
	assertDashboardEquals(t, actual, expected)
	cleanUpDashboard(t, actual.Id)
}

func TestUpdateDashboard(t *testing.T) {
	expected := getTestDashboard()
	board, err := client.CreateDashboard(expected)
	if err != nil {
		t.Fatalf("Creating a dashboard failed when it shouldn't. (%s)", err)
	}

	board.Title = "___New-Test-Board___"

	if err := client.UpdateDashboard(board); err != nil {
		t.Fatalf("Updating a dashboard failed when it shouldn't: %s", err)
	}

	actual, err := client.GetDashboard(board.Id)
	if err != nil {
		t.Fatalf("Retreiving a dashboard failed when it shouldn't: %s", err)
	}

	assertDashboardEquals(t, actual, board)
	cleanUpDashboard(t, actual.Id)
}

func TestGetDashboards(t *testing.T) {
	boards, err := client.GetDashboards()
	if err != nil {
		t.Fatalf("Retreiving dashboards failed when it shouldn't: %s", err)
	}
	num := len(boards)

	board := createTestDashboard(t)

	boards, err = client.GetDashboards()
	if err != nil {
		t.Fatalf("Retreiving dashboards failed when it shouldn't: %s", err)
	}

	if num+1 != len(boards) {
		t.Fatalf("Number of dashboards didn't match expected: %d != %d", len(boards), num+1)
	}

	cleanUpDashboard(t, board.Id)
}

func getTestDashboard() *datadog.Dashboard {
	return &datadog.Dashboard{
		Title:             "___Test-Board___",
		Description:       "Testboard description",
		TemplateVariables: []datadog.TemplateVariable{},
		Graphs:            createGraph(),
	}
}

func createTestDashboard(t *testing.T) *datadog.Dashboard {
	board := getTestDashboard()
	board, err := client.CreateDashboard(board)
	if err != nil {
		t.Fatalf("Creating a dashboard failed when it shouldn't: %s", err)
	}

	return board
}

func cleanUpDashboard(t *testing.T, id int) {
	if err := client.DeleteDashboard(id); err != nil {
		t.Fatalf("Deleting a dashboard failed when it shouldn't. (%s)", err)
	}

	deletedBoard, err := client.GetDashboard(id)
	if deletedBoard != nil {
		t.Fatal("Dashboard hasn't been deleted when it should have been.")
	}

	if err == nil {
		t.Fatal("Fetching deleted dashboard didn't lead to an error.")
	}
}

func countDashboards() int {
	boards, err := client.GetDashboards()
	if err != nil {
		log.Fatalf("Error retreiving dashboards: %s", err)
	}

	return len(boards)
}

type TestGraphDefintionRequests struct {
	Query   string `json:"q"`
	Stacked bool   `json:"stacked"`
}

func createGraph() []datadog.Graph {
	graphDefinition := datadog.Graph{}.Definition
	graphDefinition.Viz = "timeseries"
	r := datadog.Graph{}.Definition.Requests
	graphDefinition.Requests = append(r, TestGraphDefintionRequests{Query: "avg:system.mem.free{*}", Stacked: false})
	graph := datadog.Graph{Title: "Mandatory graph", Definition: graphDefinition}
	graphs := []datadog.Graph{}
	graphs = append(graphs, graph)
	return graphs
}

func assertDashboardEquals(t *testing.T, actual, expected *datadog.Dashboard) {
	if actual.Title != expected.Title {
		t.Errorf("Dashboard title does not match: %s != %s", actual.Title, expected.Title)
	}
	if actual.Description != expected.Description {
		t.Errorf("Dashboard description does not match: %s != %s", actual.Description, expected.Description)
	}
	if len(actual.Graphs) != len(expected.Graphs) {
		t.Errorf("Number of Dashboard graphs does not match: %d != %d", len(actual.Graphs), len(expected.Graphs))
	}
	if len(actual.TemplateVariables) != len(expected.TemplateVariables) {
		t.Errorf("Number of Dashboard template variables does not match: %d != %d", len(actual.TemplateVariables), len(expected.TemplateVariables))
	}
}
