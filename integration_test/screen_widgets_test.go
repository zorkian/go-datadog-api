package integration_test

import (
	"testing"

	"github.com/ojongerius/go-datadog-api"
)

func TestCheckStatusWidget(t *testing.T) {
	board := createTestScreenboard(t)

	expected := datadog.Widget{}.CheckStatusWidget

	expected.X = 1
	expected.Y = 1
	expected.Width = 5
	expected.Height = 5
	expected.TitleText = "foo"
	expected.TitleAlign = "center"
	expected.TitleSize = 1
	expected.Title = true
	expected.Type = "check_status"
	expected.Tags = "foo"
	expected.Timeframe = "1d"
	expected.Timeframe = "1d"
	expected.Check = "datadog.agent.up"
	expected.Group = "foo"
	expected.Grouping = "check"

	w := datadog.Widget{CheckStatusWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(board.Id)
	if err != nil {
		t.Fatalf("Retreiving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].CheckStatusWidget

	assertEquals(t, "x", actualWidget.X, expected.X)
	assertEquals(t, "y", actualWidget.Y, expected.Y)
	assertEquals(t, "height", actualWidget.Height, expected.Height)
	assertEquals(t, "width", actualWidget.Width, expected.Width)
	assertEquals(t, "title_text", actualWidget.TitleText, expected.TitleText)
	assertEquals(t, "title_size", actualWidget.TitleSize, expected.TitleSize)
	assertEquals(t, "title_align", actualWidget.TitleAlign, expected.TitleAlign)
	assertEquals(t, "title", actualWidget.Title, expected.Title)
	assertEquals(t, "type", actualWidget.Type, expected.Type)
	assertEquals(t, "tags", actualWidget.Tags, expected.Tags)
	assertEquals(t, "timeframe", actualWidget.Timeframe, expected.Timeframe)
	assertEquals(t, "check", actualWidget.Check, expected.Check)
	assertEquals(t, "group", actualWidget.Group, expected.Group)
	assertEquals(t, "grouping", actualWidget.Grouping, expected.Grouping)

	cleanUpScreenboard(t, board.Id)
}

func TestIFrameWidget(t *testing.T) {
	board := createTestScreenboard(t)

	expected := datadog.Widget{}.IFrameWidget

	expected.X = 1
	expected.Y = 1
	expected.Width = 5
	expected.Height = 5
	expected.TitleText = "foo"
	expected.TitleAlign = "center"
	expected.TitleSize = 1
	expected.Title = true
	expected.Url = "http://www.example.com"
	expected.Type = "iframe"

	w := datadog.Widget{IFrameWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(board.Id)
	if err != nil {
		t.Fatalf("Retreiving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].IFrameWidget

	assertEquals(t, "x", actualWidget.X, expected.X)
	assertEquals(t, "y", actualWidget.Y, expected.Y)
	assertEquals(t, "height", actualWidget.Height, expected.Height)
	assertEquals(t, "width", actualWidget.Width, expected.Width)
	assertEquals(t, "title_text", actualWidget.TitleText, expected.TitleText)
	assertEquals(t, "title_size", actualWidget.TitleSize, expected.TitleSize)
	assertEquals(t, "title_align", actualWidget.TitleAlign, expected.TitleAlign)
	assertEquals(t, "title", actualWidget.Title, expected.Title)
	assertEquals(t, "url", actualWidget.Url, expected.Url)
	assertEquals(t, "type", actualWidget.Type, expected.Type)

	cleanUpScreenboard(t, board.Id)
}

func TestNoteWidget(t *testing.T) {
	board := createTestScreenboard(t)

	expected := datadog.Widget{}.NoteWidget

	expected.X = 1
	expected.Y = 1
	expected.Width = 5
	expected.Height = 5
	expected.TitleText = "foo"
	expected.TitleAlign = "center"
	expected.TitleSize = 1
	expected.Title = true
	expected.Color = "green"
	expected.FontSize = 5
	expected.RefreshEvery = 60
	expected.TickPos = "foo"
	expected.TickEdge = "bar"
	expected.Html = "<strong>baz</strong>"
	expected.Tick = false
	expected.Note = "quz"
	expected.AutoRefresh = false

	w := datadog.Widget{NoteWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(board.Id)
	if err != nil {
		t.Fatalf("Retreiving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].NoteWidget

	assertEquals(t, "x", actualWidget.X, expected.X)
	assertEquals(t, "y", actualWidget.Y, expected.Y)
	assertEquals(t, "height", actualWidget.Height, expected.Height)
	assertEquals(t, "width", actualWidget.Width, expected.Width)
	assertEquals(t, "title_text", actualWidget.TitleText, expected.TitleText)
	assertEquals(t, "title_size", actualWidget.TitleSize, expected.TitleSize)
	assertEquals(t, "title_align", actualWidget.TitleAlign, expected.TitleAlign)
	assertEquals(t, "title", actualWidget.Title, expected.Title)
	assertEquals(t, "color", actualWidget.Color, expected.Color)
	assertEquals(t, "front_size", actualWidget.FontSize, expected.FontSize)
	assertEquals(t, "refresh_every", actualWidget.RefreshEvery, expected.RefreshEvery)
	assertEquals(t, "tick_pos", actualWidget.TickPos, expected.TickPos)
	assertEquals(t, "tick_edge", actualWidget.TickEdge, expected.TickEdge)
	assertEquals(t, "tick", actualWidget.Tick, expected.Tick)
	assertEquals(t, "html", actualWidget.Html, expected.Html)
	assertEquals(t, "note", actualWidget.Note, expected.Note)
	assertEquals(t, "auto_refresh", actualWidget.AutoRefresh, expected.AutoRefresh)

	cleanUpScreenboard(t, board.Id)
}

func TestToplistWidget(t *testing.T) {
	board := createTestScreenboard(t)

	expected := datadog.Widget{}.ToplistWidget
	expected.X = 1
	expected.Y = 1
	expected.Width = 5
	expected.Height = 5
	expected.Type = "toplist"
	expected.TitleText = "foo"
	expected.TitleSize.Auto = false
	expected.TitleSize.Size = 5
	expected.TitleAlign = "center"
	expected.Title = false
	expected.Timeframe = "5m"
	expected.Legend = false
	expected.LegendSize = "5"

	w := datadog.Widget{ToplistWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(board.Id)
	if err != nil {
		t.Fatalf("Retreiving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].ToplistWidget

	assertEquals(t, "x", actualWidget.X, expected.X)
	assertEquals(t, "y", actualWidget.Y, expected.Y)
	assertEquals(t, "height", actualWidget.Height, expected.Height)
	assertEquals(t, "width", actualWidget.Width, expected.Width)
	assertEquals(t, "title_text", actualWidget.TitleText, expected.TitleText)
	assertEquals(t, "title_size", actualWidget.TitleSize, expected.TitleSize)
	assertEquals(t, "title_align", actualWidget.TitleAlign, expected.TitleAlign)
	assertEquals(t, "legend", actualWidget.Legend, expected.Legend)
	assertEquals(t, "legend_size", actualWidget.LegendSize, expected.LegendSize)

	cleanUpScreenboard(t, board.Id)
}

func TestEventSteamWidget(t *testing.T) {
	board := createTestScreenboard(t)

	expected := datadog.Widget{}.EventStreamWidget
	expected.EventSize = "1"
	expected.Width = 1
	expected.Height = 1
	expected.X = 1
	expected.Y = 1
	expected.Query = "foo"
	expected.Timeframe = "5w"
	expected.Title = false
	expected.TitleAlign = "center"
	expected.TitleSize.Auto = false
	expected.TitleSize.Size = 5
	expected.TitleText = "bar"
	expected.Type = "baz"

	w := datadog.Widget{EventStreamWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(board.Id)
	if err != nil {
		t.Fatalf("Retreiving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].EventStreamWidget

	assertEquals(t, "event_size", actualWidget.EventSize, expected.EventSize)
	assertEquals(t, "width", actualWidget.Width, expected.Width)
	assertEquals(t, "height", actualWidget.Height, expected.Height)
	assertEquals(t, "x", actualWidget.X, expected.X)
	assertEquals(t, "y", actualWidget.Y, expected.Y)
	assertEquals(t, "query", actualWidget.Query, expected.Query)
	assertEquals(t, "timeframe", actualWidget.Timeframe, expected.Timeframe)
	assertEquals(t, "title", actualWidget.Title, expected.Title)
	assertEquals(t, "title_align", actualWidget.TitleAlign, expected.TitleAlign)
	assertEquals(t, "title_size", actualWidget.TitleSize, expected.TitleSize)
	assertEquals(t, "title_text", actualWidget.TitleText, expected.TitleText)
	assertEquals(t, "type", actualWidget.Type, expected.Type)

	cleanUpScreenboard(t, board.Id)
}

func TestImageWidget(t *testing.T) {
	board := createTestScreenboard(t)

	expected := datadog.Widget{}.ImageWidget

	expected.Width = 1
	expected.Height = 1
	expected.X = 1
	expected.Y = 1
	expected.Title = false
	expected.TitleAlign = "center"
	expected.TitleSize.Auto = false
	expected.TitleSize.Size = 5
	expected.TitleText = "bar"
	expected.Type = "baz"
	expected.Url = "qux"
	expected.Sizing = "quuz"

	w := datadog.Widget{ImageWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(board.Id)
	if err != nil {
		t.Fatalf("Retreiving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].ImageWidget

	assertEquals(t, "width", actualWidget.Width, expected.Width)
	assertEquals(t, "height", actualWidget.Height, expected.Height)
	assertEquals(t, "x", actualWidget.X, expected.X)
	assertEquals(t, "y", actualWidget.Y, expected.Y)
	assertEquals(t, "title", actualWidget.Title, expected.Title)
	assertEquals(t, "title_align", actualWidget.TitleAlign, expected.TitleAlign)
	assertEquals(t, "title_size", actualWidget.TitleSize, expected.TitleSize)
	assertEquals(t, "title_text", actualWidget.TitleText, expected.TitleText)
	assertEquals(t, "type", actualWidget.Type, expected.Type)
	assertEquals(t, "url", actualWidget.Url, expected.Url)
	assertEquals(t, "sizing", actualWidget.Sizing, expected.Sizing)

	cleanUpScreenboard(t, board.Id)
}

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

	w := datadog.Widget{FreeTextWidget: expected}

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

	w := datadog.Widget{TimeseriesWidget: expected}

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

	w := datadog.Widget{QueryValueWidget: expected}

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
