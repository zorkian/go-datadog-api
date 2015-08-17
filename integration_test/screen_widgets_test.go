package integration_test

import (
	"testing"

	"github.com/ojongerius/go-datadog-api"
)

/* TODO: Add tests for:
	* ToplistWidget
	* EventStreamWidget
	* ImageWidget
*/

func TestFreeTextWidget(t *testing.T) {
	board := createTestScreenboard(t)

	expected := datadog.Widget{}.FreeTextWidget

	expected.X = 1
	expected.Y = 1
	expected.Height = 10
	expected.Width = 10
	expected.Text = "Test"
	expected.FontSize = "16"
	expected.TextAlign = "center"

	w := datadog.Widget{ FreeTextWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(board.Id)
	if err != nil {
		t.Fatalf("Retreiving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].FreeTextWidget

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

	expected := datadog.Widget{}.TimeseriesWidget
	expected.X = 1
	expected.Y = 1
	expected.Width = 20
	expected.Height = 30
	expected.Title = true
	expected.TitleAlign = "centre"
	expected.TitleSize = datadog.TextSize{Size: 16}
	expected.TitleText = "Test"
	expected.Timeframe = "1m"

	w := datadog.Widget{ TimeseriesWidget: expected}

	board.Widgets = append(board.Widgets, w)
	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(board.Id)
	if err != nil {
		t.Fatalf("Retreiving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].TimeseriesWidget

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

	expected := datadog.Widget{}.QueryValueWidget
	expected.X = 1
	expected.Y = 1
	expected.Width = 20
	expected.Height = 30
	expected.Title = true
	expected.TitleAlign = "centre"
	expected.TitleSize = datadog.TextSize{Size: 16}
	expected.TitleText = "Test"
	expected.Timeframe = "1m"
	expected.TimeframeAggregator = "sum"
	expected.Aggregator = "min"
	expected.Query = "docker.containers.running"

	w := datadog.Widget{ QueryValueWidget: expected}

	board.Widgets = append(board.Widgets, w)
	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(board.Id)
	if err != nil {
		t.Fatalf("Retreiving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].QueryValueWidget

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
