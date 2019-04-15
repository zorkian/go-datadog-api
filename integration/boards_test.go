package integration

import (
	"testing"

	"github.com/zorkian/go-datadog-api"
)

func TestOrderedBoardCreateAndDelete(t *testing.T) {
	expected := getTestBoard()

	// Create the dashboard and compare it
	actual, err := client.CreateBoard(expected)
	if err != nil {
		t.Fatalf("Creating a dashboard failed when it shouldn't. (%s)", err)
	}
	defer cleanUpBoard(t, *actual.Id)
	assertBoardEquals(t, actual, expected)

	// Now try to fetch it freshly and compare it again
	actual, err = client.GetBoard(*actual.Id)
	if err != nil {
		t.Fatalf("Retrieving a dashboard failed when it shouldn't. (%s)", err)
	}
	assertBoardEquals(t, actual, expected)
}

// Helpers for the tests

func getTestBoard() *datadog.Board {
	return &datadog.Board{
		Title:       datadog.String("Test Dashboard"),
		Widgets:     createWidgets(),
		LayoutType:  datadog.String("ordered"),
		Description: datadog.String("Test Dashboard description"),
		IsReadOnly:  datadog.Bool(false),
	}
}

func createWidgets() []datadog.BoardWidget {
	widgets := []datadog.BoardWidget{}
	widgets = append(widgets, createAlertGraph())
	widgets = append(widgets, createTimeseriesWidget())
	widgets = append(widgets, createGroupWidget())
	return widgets
}

func createAlertGraph() datadog.BoardWidget {
	return datadog.BoardWidget{
		Definition: &datadog.AlertGraphDefinition{
			Type:    datadog.String("alert_graph"),
			AlertId: datadog.String("123456"),
			VizType: datadog.String("timeseries"),
			Title:   datadog.String("Test Alert Graph widget"),
		},
	}
}

func createGroupWidget() datadog.BoardWidget {
	widgets := []datadog.BoardWidget{
		createTimeseriesWidget(),
		createAlertGraph(),
	}
	return datadog.BoardWidget{
		Definition: &datadog.GroupDefinition{
			Type:       datadog.String("group"),
			LayoutType: datadog.String("ordered"),
			Widgets:    widgets,
			Title:      datadog.String("Group widget"),
		},
	}
}

func createTimeseriesWidget() datadog.BoardWidget {
	request := datadog.TimeseriesRequest{
		MetricQuery: datadog.String("avg:system.cpu.user{*}"),
		Style: &datadog.TimeseriesRequestStyle{
			LineType: datadog.String("dotted"),
		},
		DisplayType: datadog.String("area"),
	}
	return datadog.BoardWidget{
		Definition: &datadog.TimeseriesDefinition{
			Type:     datadog.String("timeseries"),
			Requests: []datadog.TimeseriesRequest{request},
			Yaxis: &datadog.WidgetAxis{
				Label:       datadog.String("y-axis label"),
				Max:         datadog.String("3000"),
				IncludeZero: datadog.Bool(true),
			},
			Title: datadog.String("Test Timeseries widget"),
		},
	}
}

func cleanUpBoard(t *testing.T, id string) {
	if err := client.DeleteBoard(id); err != nil {
		t.Fatalf("Deleting a dashboard failed when it shouldn't. Manual cleanup needed. (%s)", err)
	}

	deletedBoard, err := client.GetBoard(id)
	if deletedBoard != nil {
		t.Fatal("Dashboard hasn't been deleted when it should have been. Manual cleanup needed.")
	}

	if err == nil {
		t.Fatal("Fetching deleted dashboard didn't lead to an error. Manual cleanup needed.")
	}
}

func assertBoardEquals(t *testing.T, actual, expected *datadog.Board) {
	if *actual.Title != *expected.Title {
		t.Errorf("Dashboard title does not match: %s != %s", *actual.Title, *expected.Title)
	}
	if len(actual.Widgets) != len(expected.Widgets) {
		t.Errorf("Number of Dashboard widgets does not match: %d != %d", len(actual.Widgets), len(expected.Widgets))
	}
	if *actual.LayoutType != *expected.LayoutType {
		t.Errorf("Dashboard layout type does not match: %s != %s", *actual.LayoutType, *expected.LayoutType)
	}
	if *actual.Description != *expected.Description {
		t.Errorf("Dashboard description does not match: %s != %s", *actual.Description, *expected.Description)
	}
	if len(actual.TemplateVariables) != len(expected.TemplateVariables) {
		t.Errorf("Number of Dashboard template variables does not match: %d != %d", len(actual.TemplateVariables), len(expected.TemplateVariables))
	}
	if *actual.IsReadOnly != *expected.IsReadOnly {
		t.Errorf("Dashboard description does not match: %v != %v", *actual.IsReadOnly, *expected.IsReadOnly)
	}
	if len(actual.NotifyList) != len(expected.NotifyList) {
		t.Errorf("Number of users to notify does not match: %d != %d", len(actual.NotifyList), len(expected.NotifyList))
	}
}
