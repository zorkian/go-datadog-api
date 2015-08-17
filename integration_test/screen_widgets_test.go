package integration_test

import (
	"testing"

	"github.com/seiffert/go-datadog-api"
)

func TestFreeTextWidget(t *testing.T) {
	board := createTestScreenboard(t)
	widget := datadog.NewFreeTextWidget(
		1, 1, 10, 10, "Test", 16, "center",
	)
	expected := *(widget.(*datadog.FreeTextWidget))

	board.Widgets = append(board.Widgets, widget)
	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(board.Id)
	if err != nil {
		t.Fatalf("Retreiving a screenboard failed: %s", err)
	}

	actualWidget, ok := actual.Widgets[0].(*datadog.FreeTextWidget)
	if !ok {
		t.Fatalf("Widget type does not match: %v", actual.Widgets[0])
	}

	assertEquals(t, "font-size", actualWidget.FontSize, expected.FontSize)
	assertEquals(t, "height", actualWidget.Height, expected.Height)
	assertEquals(t, "width", actualWidget.Width, expected.Width)
	assertEquals(t, "x", actualWidget.X, expected.X)
	assertEquals(t, "y", actualWidget.Y, expected.Y)
	assertEquals(t, "text", actualWidget.Text, expected.Text)
	assertEquals(t, "text-align", actualWidget.TextAlign, expected.TextAlign)
	assertEquals(t, "type", actualWidget.Type, expected.Type)

	cleanUpScreenboard(t, board.Id)
}

func TestTimeseriesWidget(t *testing.T) {
	board := createTestScreenboard(t)
	widget := datadog.NewTimeseriesWidget(
		1, 1, 20, 30,
		true, "center", datadog.TextSize{Size: 16}, "Test",
		"1m",
		[]datadog.TimeseriesRequest{
			datadog.NewTimeseriesRequest("line", "system.cpu.idle"),
		},
	)
	expected := *(widget.(*datadog.TimeseriesWidget))

	board.Widgets = append(board.Widgets, widget)
	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(board.Id)
	if err != nil {
		t.Fatalf("Retreiving a screenboard failed: %s", err)
	}

	actualWidget, ok := actual.Widgets[0].(*datadog.TimeseriesWidget)
	if !ok {
		t.Fatalf("Widget type does not match: %v", actual.Widgets[0])
	}

	assertEquals(t, "height", actualWidget.Height, expected.Height)
	assertEquals(t, "width", actualWidget.Width, expected.Width)
	assertEquals(t, "x", actualWidget.X, expected.X)
	assertEquals(t, "y", actualWidget.Y, expected.Y)
	assertEquals(t, "title", actualWidget.Title, expected.Title)
	assertEquals(t, "title-align", actualWidget.TitleAlign, expected.TitleAlign)
	assertEquals(t, "title-size.size", actualWidget.TitleSize.Size, expected.TitleSize.Size)
	assertEquals(t, "title-size.auto", actualWidget.TitleSize.Auto, expected.TitleSize.Auto)
	assertEquals(t, "title-text", actualWidget.TitleText, expected.TitleText)
	assertEquals(t, "type", actualWidget.Type, expected.Type)
	assertEquals(t, "timeframe", actualWidget.Timeframe, expected.Timeframe)
	assertEquals(t, "legend", actualWidget.Legend, expected.Legend)
	assertTileDefEquals(t, actualWidget.TileDef, expected.TileDef)

	cleanUpScreenboard(t, board.Id)
}

func TestQueryValueWidget(t *testing.T) {
	board := createTestScreenboard(t)
	widget := datadog.NewQueryValueWidget(
		1, 1, 20, 10,
		true, "center", datadog.TextSize{Size: 16}, "Test",
		"left", datadog.TextSize{Size: 32},
		"1m", "sum",
		"min", "docker.containers.running",
	)
	expected := *(widget.(*datadog.QueryValueWidget))

	board.Widgets = append(board.Widgets, widget)
	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(board.Id)
	if err != nil {
		t.Fatalf("Retreiving a screenboard failed: %s", err)
	}

	actualWidget, ok := actual.Widgets[0].(*datadog.QueryValueWidget)
	if !ok {
		t.Fatalf("Widget type does not match: %v", actual.Widgets[0])
	}

	assertEquals(t, "height", actualWidget.Height, expected.Height)
	assertEquals(t, "width", actualWidget.Width, expected.Width)
	assertEquals(t, "x", actualWidget.X, expected.X)
	assertEquals(t, "y", actualWidget.Y, expected.Y)
	assertEquals(t, "title", actualWidget.Title, expected.Title)
	assertEquals(t, "title-align", actualWidget.TitleAlign, expected.TitleAlign)
	assertEquals(t, "title-size.size", actualWidget.TitleSize.Size, expected.TitleSize.Size)
	assertEquals(t, "title-size.auto", actualWidget.TitleSize.Auto, expected.TitleSize.Auto)
	assertEquals(t, "title-text", actualWidget.TitleText, expected.TitleText)
	assertEquals(t, "type", actualWidget.Type, expected.Type)
	assertEquals(t, "timeframe", actualWidget.Timeframe, expected.Timeframe)
	assertEquals(t, "timeframe-aggregator", actualWidget.TimeframeAggregator, expected.TimeframeAggregator)
	assertEquals(t, "aggregator", actualWidget.Aggregator, expected.Aggregator)
	assertEquals(t, "query", actualWidget.Query, expected.Query)

	cleanUpScreenboard(t, board.Id)
}

func assertTileDefEquals(t *testing.T, actual datadog.TileDef, expected datadog.TileDef) {
	assertEquals(t, "num-events", len(actual.Events), len(expected.Events))
	assertEquals(t, "num-requests", len(actual.Requests), len(expected.Requests))
	assertEquals(t, "viz", actual.Viz, expected.Viz)

	for i, event := range actual.Events {
		assertEquals(t, "event-query", event.Query, expected.Events[i].Query)
	}

	for i, request := range actual.Requests {
		assertEquals(t, "request-query", request.Query, expected.Requests[i].Query)
		assertEquals(t, "request-type", request.Type, expected.Requests[i].Type)
	}
}

func assertEquals(t *testing.T, attribute string, a, b interface{}) {
	if a != b {
		t.Errorf("The two %s values '%v' and '%v' are not equal", attribute, a, b)
	}
}
