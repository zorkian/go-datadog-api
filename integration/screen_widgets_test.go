package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
)

func TestAlertValueWidget(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.AlertValueWidget{}

	expected.X = datadog.Int(1)
	expected.Y = datadog.Int(1)
	expected.Width = datadog.Int(5)
	expected.Height = datadog.Int(5)
	expected.TitleText = datadog.String("foo")
	expected.TitleAlign = datadog.String("center")
	expected.TitleSize = datadog.Int(1)
	expected.Title = datadog.Bool(true)
	expected.TextSize = datadog.String("auto")
	expected.Precision = datadog.Int(2)
	expected.AlertId = datadog.Int(1)
	expected.Type = datadog.String("alert_value")
	expected.Unit = datadog.String("auto")
	expected.AddTimeframe = datadog.Bool(false)

	w := datadog.Widget{AlertValueWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].AlertValueWidget

	assert.Equal(t, *actualWidget.X, *expected.X)
	assert.Equal(t, *actualWidget.Y, *expected.Y)
	assert.Equal(t, *actualWidget.Height, *expected.Height)
	assert.Equal(t, *actualWidget.Width, *expected.Width)
	assert.Equal(t, *actualWidget.TitleText, *expected.TitleText)
	assert.Equal(t, *actualWidget.TitleSize, *expected.TitleSize)
	assert.Equal(t, *actualWidget.TitleAlign, *expected.TitleAlign)
	assert.Equal(t, *actualWidget.Title, *expected.Title)

	assert.Equal(t, *actualWidget.TextSize, *expected.TextSize)
	assert.Equal(t, *actualWidget.Precision, *expected.Precision)
	assert.Equal(t, *actualWidget.AlertId, *expected.AlertId)
	assert.Equal(t, *actualWidget.Type, *expected.Type)
	assert.Equal(t, *actualWidget.Unit, *expected.Unit)
	assert.Equal(t, *actualWidget.AddTimeframe, *expected.AddTimeframe)
}

func TestChangeWidget(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.ChangeWidget{}

	expected.X = datadog.Int(1)
	expected.Y = datadog.Int(1)
	expected.Width = datadog.Int(5)
	expected.Height = datadog.Int(5)
	expected.TitleText = datadog.String("foo")
	expected.TitleAlign = datadog.String("center")
	expected.TitleSize = datadog.Int(1)
	expected.Title = datadog.Bool(true)

	expected.Aggregator = datadog.String("min")
	expected.TileDef = &datadog.TileDef{}

	w := datadog.Widget{ChangeWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].ChangeWidget

	assert.Equal(t, *actualWidget.X, *expected.X)
	assert.Equal(t, *actualWidget.Y, *expected.Y)
	assert.Equal(t, *actualWidget.Height, *expected.Height)
	assert.Equal(t, *actualWidget.Width, *expected.Width)
	assert.Equal(t, *actualWidget.TitleText, *expected.TitleText)
	assert.Equal(t, *actualWidget.TitleSize, *expected.TitleSize)
	assert.Equal(t, *actualWidget.TitleAlign, *expected.TitleAlign)
	assert.Equal(t, *actualWidget.Title, *expected.Title)

	assert.Equal(t, *actualWidget.Aggregator, *expected.Aggregator)
	assertTileDefEquals(t, *actualWidget.TileDef, *expected.TileDef)
}

func TestGraphWidget(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.GraphWidget{}

	expected.X = datadog.Int(1)
	expected.Y = datadog.Int(1)
	expected.Width = datadog.Int(5)
	expected.Height = datadog.Int(5)
	expected.TitleText = datadog.String("foo")
	expected.TitleAlign = datadog.String("center")
	expected.TitleSize = datadog.Int(1)
	expected.Title = datadog.Bool(true)

	expected.Timeframe = datadog.String("1d")
	expected.Type = datadog.String("alert_graph")
	expected.Legend = datadog.Bool(true)
	expected.LegendSize = datadog.Int(5)
	expected.TileDef = &datadog.TileDef{}

	w := datadog.Widget{GraphWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].GraphWidget

	assert.Equal(t, *actualWidget.X, *expected.X)
	assert.Equal(t, *actualWidget.Y, *expected.Y)
	assert.Equal(t, *actualWidget.Height, *expected.Height)
	assert.Equal(t, *actualWidget.Width, *expected.Width)
	assert.Equal(t, *actualWidget.TitleText, *expected.TitleText)
	assert.Equal(t, *actualWidget.TitleSize, *expected.TitleSize)
	assert.Equal(t, *actualWidget.TitleAlign, *expected.TitleAlign)
	assert.Equal(t, *actualWidget.Title, *expected.Title)

	assert.Equal(t, *actualWidget.Type, *expected.Type)
	assert.Equal(t, *actualWidget.Timeframe, *expected.Timeframe)
	assert.Equal(t, *actualWidget.Legend, *expected.Legend)
	assert.Equal(t, *actualWidget.LegendSize, *expected.LegendSize)
	assertTileDefEquals(t, *actualWidget.TileDef, *expected.TileDef)
}

func TestEventTimelineWidget(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.EventTimelineWidget{}

	expected.X = datadog.Int(1)
	expected.Y = datadog.Int(1)
	expected.Width = datadog.Int(5)
	expected.Height = datadog.Int(5)
	expected.TitleText = datadog.String("foo")
	expected.TitleAlign = datadog.String("center")
	expected.TitleSize = datadog.Int(1)
	expected.Title = datadog.Bool(true)

	expected.Query = datadog.String("avg:system.load.1{foo} by {bar}")
	expected.Timeframe = datadog.String("1d")
	expected.Type = datadog.String("alert_graph")

	w := datadog.Widget{EventTimelineWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].EventTimelineWidget

	assert.Equal(t, *actualWidget.X, *expected.X)
	assert.Equal(t, *actualWidget.Y, *expected.Y)
	assert.Equal(t, *actualWidget.Height, *expected.Height)
	assert.Equal(t, *actualWidget.Width, *expected.Width)
	assert.Equal(t, *actualWidget.TitleText, *expected.TitleText)
	assert.Equal(t, *actualWidget.TitleSize, *expected.TitleSize)
	assert.Equal(t, *actualWidget.TitleAlign, *expected.TitleAlign)
	assert.Equal(t, *actualWidget.Title, *expected.Title)

	assert.Equal(t, *actualWidget.Type, *expected.Type)
	assert.Equal(t, *actualWidget.Query, *expected.Query)
	assert.Equal(t, *actualWidget.Timeframe, *expected.Timeframe)
}

func TestAlertGraphWidget(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.AlertGraphWidget{}

	expected.X = datadog.Int(1)
	expected.Y = datadog.Int(1)
	expected.Width = datadog.Int(5)
	expected.Height = datadog.Int(5)
	expected.TitleText = datadog.String("foo")
	expected.TitleAlign = datadog.String("center")
	expected.TitleSize = datadog.Int(1)
	expected.Title = datadog.Bool(true)

	expected.VizType = datadog.String("")
	expected.Timeframe = datadog.String("1d")
	expected.AddTimeframe = datadog.Bool(false)
	expected.AlertId = datadog.Int(1)
	expected.Type = datadog.String("alert_graph")

	w := datadog.Widget{AlertGraphWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].AlertGraphWidget

	assert.Equal(t, *actualWidget.X, *expected.X)
	assert.Equal(t, *actualWidget.Y, *expected.Y)
	assert.Equal(t, *actualWidget.Height, *expected.Height)
	assert.Equal(t, *actualWidget.Width, *expected.Width)
	assert.Equal(t, *actualWidget.TitleText, *expected.TitleText)
	assert.Equal(t, *actualWidget.TitleSize, *expected.TitleSize)
	assert.Equal(t, *actualWidget.TitleAlign, *expected.TitleAlign)
	assert.Equal(t, *actualWidget.Title, *expected.Title)

	assert.Equal(t, *actualWidget.Type, *expected.Type)
	assert.Equal(t, *actualWidget.VizType, *expected.VizType)
	assert.Equal(t, *actualWidget.Timeframe, *expected.Timeframe)
	assert.Equal(t, *actualWidget.AddTimeframe, *expected.AddTimeframe)
	assert.Equal(t, *actualWidget.AlertId, *expected.AlertId)
}

func TestHostMapWidget(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.HostMapWidget{}

	expected.X = datadog.Int(1)
	expected.Y = datadog.Int(1)
	expected.Width = datadog.Int(5)
	expected.Height = datadog.Int(5)
	expected.TitleText = datadog.String("foo")
	expected.TitleAlign = datadog.String("center")
	expected.TitleSize = datadog.Int(1)
	expected.Title = datadog.Bool(true)

	expected.Type = datadog.String("check_status")
	expected.Query = datadog.String("avg:system.load.1{foo} by {bar}")
	expected.Timeframe = datadog.String("1d")
	expected.Legend = datadog.Bool(true)
	expected.LegendSize = datadog.Int(5)
	expected.TileDef = &datadog.TileDef{}

	w := datadog.Widget{HostMapWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].HostMapWidget

	assert.Equal(t, *actualWidget.X, *expected.X)
	assert.Equal(t, *actualWidget.Y, *expected.Y)
	assert.Equal(t, *actualWidget.Height, *expected.Height)
	assert.Equal(t, *actualWidget.Width, *expected.Width)
	assert.Equal(t, *actualWidget.TitleText, *expected.TitleText)
	assert.Equal(t, *actualWidget.TitleSize, *expected.TitleSize)
	assert.Equal(t, *actualWidget.TitleAlign, *expected.TitleAlign)
	assert.Equal(t, *actualWidget.Title, *expected.Title)

	assert.Equal(t, *actualWidget.Type, *expected.Type)
	assert.Equal(t, *actualWidget.Query, *expected.Query)
	assert.Equal(t, *actualWidget.Timeframe, *expected.Timeframe)
	assert.Equal(t, *actualWidget.Query, *expected.Query)
	assert.Equal(t, *actualWidget.Legend, *expected.Legend)
	assert.Equal(t, *actualWidget.LegendSize, *expected.LegendSize)
	assertTileDefEquals(t, *actualWidget.TileDef, *expected.TileDef)
}

func TestCheckStatusWidget(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.CheckStatusWidget{}

	expected.X = datadog.Int(1)
	expected.Y = datadog.Int(1)
	expected.Width = datadog.Int(5)
	expected.Height = datadog.Int(5)
	expected.TitleText = datadog.String("foo")
	expected.TitleAlign = datadog.String("center")
	expected.TitleSize = datadog.Int(1)
	expected.Title = datadog.Bool(true)

	expected.Type = datadog.String("check_status")
	expected.Tags = datadog.String("foo")
	expected.Timeframe = datadog.String("1d")
	expected.Timeframe = datadog.String("1d")
	expected.Check = datadog.String("datadog.agent.up")
	expected.Group = datadog.String("foo")
	expected.Grouping = datadog.String("check")

	w := datadog.Widget{CheckStatusWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].CheckStatusWidget

	assert.Equal(t, *actualWidget.X, *expected.X)
	assert.Equal(t, *actualWidget.Y, *expected.Y)
	assert.Equal(t, *actualWidget.Height, *expected.Height)
	assert.Equal(t, *actualWidget.Width, *expected.Width)
	assert.Equal(t, *actualWidget.TitleText, *expected.TitleText)
	assert.Equal(t, *actualWidget.TitleSize, *expected.TitleSize)
	assert.Equal(t, *actualWidget.TitleAlign, *expected.TitleAlign)
	assert.Equal(t, *actualWidget.Title, *expected.Title)

	assert.Equal(t, *actualWidget.Type, *expected.Type)
	assert.Equal(t, *actualWidget.Tags, *expected.Tags)
	assert.Equal(t, *actualWidget.Timeframe, *expected.Timeframe)
	assert.Equal(t, *actualWidget.Check, *expected.Check)
	assert.Equal(t, *actualWidget.Group, *expected.Group)
	assert.Equal(t, *actualWidget.Grouping, *expected.Grouping)
}

func TestIFrameWidget(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.IFrameWidget{}

	expected.X = datadog.Int(1)
	expected.Y = datadog.Int(1)
	expected.Width = datadog.Int(5)
	expected.Height = datadog.Int(5)
	expected.TitleText = datadog.String("foo")
	expected.TitleAlign = datadog.String("center")
	expected.TitleSize = datadog.Int(1)
	expected.Title = datadog.Bool(true)

	expected.Url = datadog.String("http://www.example.com")
	expected.Type = datadog.String("iframe")

	w := datadog.Widget{IFrameWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].IFrameWidget

	assert.Equal(t, *actualWidget.X, *expected.X)
	assert.Equal(t, *actualWidget.Y, *expected.Y)
	assert.Equal(t, *actualWidget.Height, *expected.Height)
	assert.Equal(t, *actualWidget.Width, *expected.Width)
	assert.Equal(t, *actualWidget.TitleText, *expected.TitleText)
	assert.Equal(t, *actualWidget.TitleSize, *expected.TitleSize)
	assert.Equal(t, *actualWidget.TitleAlign, *expected.TitleAlign)
	assert.Equal(t, *actualWidget.Title, *expected.Title)

	assert.Equal(t, *actualWidget.Url, *expected.Url)
	assert.Equal(t, *actualWidget.Type, *expected.Type)
}

func TestNoteWidget(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.NoteWidget{}

	expected.X = datadog.Int(1)
	expected.Y = datadog.Int(1)
	expected.Width = datadog.Int(5)
	expected.Height = datadog.Int(5)
	expected.TitleText = datadog.String("foo")
	expected.TitleAlign = datadog.String("center")
	expected.TitleSize = datadog.Int(1)
	expected.Title = datadog.Bool(true)

	expected.Color = datadog.String("green")
	expected.FontSize = datadog.Int(5)
	expected.RefreshEvery = datadog.Int(60)
	expected.TickPos = datadog.String("foo")
	expected.TickEdge = datadog.String("bar")
	expected.Html = datadog.String("<strong>baz</strong>")
	expected.Tick = datadog.Bool(false)
	expected.Note = datadog.String("quz")
	expected.AutoRefresh = datadog.Bool(false)

	w := datadog.Widget{NoteWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].NoteWidget

	assert.Equal(t, *actualWidget.X, *expected.X)
	assert.Equal(t, *actualWidget.Y, *expected.Y)
	assert.Equal(t, *actualWidget.Height, *expected.Height)
	assert.Equal(t, *actualWidget.Width, *expected.Width)
	assert.Equal(t, *actualWidget.TitleText, *expected.TitleText)
	assert.Equal(t, *actualWidget.TitleSize, *expected.TitleSize)
	assert.Equal(t, *actualWidget.TitleAlign, *expected.TitleAlign)
	assert.Equal(t, *actualWidget.Title, *expected.Title)

	assert.Equal(t, *actualWidget.Color, *expected.Color)
	assert.Equal(t, *actualWidget.FontSize, *expected.FontSize)
	assert.Equal(t, *actualWidget.RefreshEvery, *expected.RefreshEvery)
	assert.Equal(t, *actualWidget.TickPos, *expected.TickPos)
	assert.Equal(t, *actualWidget.TickEdge, *expected.TickEdge)
	assert.Equal(t, *actualWidget.Tick, *expected.Tick)
	assert.Equal(t, *actualWidget.Html, *expected.Html)
	assert.Equal(t, *actualWidget.Note, *expected.Note)
	assert.Equal(t, *actualWidget.AutoRefresh, *expected.AutoRefresh)
}

func TestToplistWidget(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.ToplistWidget{}

	expected.X = datadog.Int(1)
	expected.Y = datadog.Int(1)
	expected.Width = datadog.Int(5)
	expected.Height = datadog.Int(5)
	expected.TitleText = datadog.String("foo")
	expected.TitleAlign = datadog.String("center")
	expected.TitleSize = &datadog.TextSize{
		Size: datadog.Int(1),
		Auto: datadog.Bool(true),
	}
	expected.Title = datadog.Bool(true)
	expected.TitleAlign = datadog.String("center")
	expected.Timeframe = datadog.String("5m")
	expected.Legend = datadog.Bool(false)
	expected.LegendSize = datadog.Int(5)

	w := datadog.Widget{ToplistWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].ToplistWidget

	assert.Equal(t, *actualWidget.X, *expected.X)
	assert.Equal(t, *actualWidget.Y, *expected.Y)
	assert.Equal(t, *actualWidget.Height, *expected.Height)
	assert.Equal(t, *actualWidget.Width, *expected.Width)
	assert.Equal(t, *actualWidget.TitleText, *expected.TitleText)
	assert.Equal(t, *actualWidget.TitleSize, *expected.TitleSize)
	assert.Equal(t, *actualWidget.TitleAlign, *expected.TitleAlign)
	assert.Equal(t, *actualWidget.Title, *expected.Title)

	assert.Equal(t, *actualWidget.Legend, *expected.Legend)
	assert.Equal(t, *actualWidget.LegendSize, *expected.LegendSize)
}

func TestEventSteamWidget(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.EventStreamWidget{}

	expected.EventSize = datadog.String("1")
	expected.X = datadog.Int(1)
	expected.Y = datadog.Int(1)
	expected.Width = datadog.Int(5)
	expected.Height = datadog.Int(5)
	expected.TitleText = datadog.String("foo")
	expected.TitleAlign = datadog.String("center")
	expected.TitleSize = &datadog.TextSize{
		Size: datadog.Int(1),
		Auto: datadog.Bool(true),
	}
	expected.Title = datadog.Bool(true)
	expected.TitleAlign = datadog.String("center")
	expected.Timeframe = datadog.String("5m")

	expected.Query = datadog.String("foo")
	expected.Timeframe = datadog.String("5w")
	expected.Title = datadog.Bool(false)
	expected.TitleAlign = datadog.String("center")
	expected.TitleSize.Auto = datadog.Bool(false)
	expected.TitleSize.Size = datadog.Int(5)
	expected.TitleText = datadog.String("bar")
	expected.Type = datadog.String("baz")

	w := datadog.Widget{EventStreamWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].EventStreamWidget

	assert.Equal(t, *actualWidget.X, *expected.X)
	assert.Equal(t, *actualWidget.Y, *expected.Y)
	assert.Equal(t, *actualWidget.Height, *expected.Height)
	assert.Equal(t, *actualWidget.Width, *expected.Width)
	assert.Equal(t, *actualWidget.TitleText, *expected.TitleText)
	assert.Equal(t, *actualWidget.TitleSize, *expected.TitleSize)
	assert.Equal(t, *actualWidget.TitleAlign, *expected.TitleAlign)
	assert.Equal(t, *actualWidget.Title, *expected.Title)

	assert.Equal(t, *actualWidget.Type, *expected.Type)
}

func TestImageWidget(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.ImageWidget{}

	expected.X = datadog.Int(1)
	expected.Y = datadog.Int(1)
	expected.Width = datadog.Int(5)
	expected.Height = datadog.Int(5)
	expected.Title = datadog.Bool(false)
	expected.TitleAlign = datadog.String("center")
	expected.TitleSize = &datadog.TextSize{
		Size: datadog.Int(1),
		Auto: datadog.Bool(true),
	}
	expected.TitleText = datadog.String("bar")
	expected.Type = datadog.String("baz")
	expected.Url = datadog.String("qux")
	expected.Sizing = datadog.String("quuz")

	w := datadog.Widget{ImageWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].ImageWidget

	assert.Equal(t, *actualWidget.X, *expected.X)
	assert.Equal(t, *actualWidget.Y, *expected.Y)
	assert.Equal(t, *actualWidget.Height, *expected.Height)
	assert.Equal(t, *actualWidget.Width, *expected.Width)
	assert.Equal(t, *actualWidget.TitleText, *expected.TitleText)
	assert.Equal(t, *actualWidget.TitleSize, *expected.TitleSize)
	assert.Equal(t, *actualWidget.TitleAlign, *expected.TitleAlign)
	assert.Equal(t, *actualWidget.Title, *expected.Title)

	assert.Equal(t, *actualWidget.Type, *expected.Type)
	assert.Equal(t, *actualWidget.Url, *expected.Url)
	assert.Equal(t, *actualWidget.Sizing, *expected.Sizing)
}

func TestFreeTextWidget(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.FreeTextWidget{}

	expected.X = datadog.Int(1)
	expected.Y = datadog.Int(1)
	expected.Width = datadog.Int(5)
	expected.Height = datadog.Int(5)
	expected.Text = datadog.String("Test")
	expected.FontSize = datadog.String("16")
	expected.TextAlign = datadog.String("center")

	w := datadog.Widget{FreeTextWidget: expected}

	board.Widgets = append(board.Widgets, w)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].FreeTextWidget

	assert.Equal(t, *actualWidget.X, *expected.X)
	assert.Equal(t, *actualWidget.Y, *expected.Y)
	assert.Equal(t, *actualWidget.Height, *expected.Height)
	assert.Equal(t, *actualWidget.Width, *expected.Width)

	assert.Equal(t, "font-size", *actualWidget.FontSize, *expected.FontSize)
	assert.Equal(t, "type", *actualWidget.Type, *expected.Type)
}

func TestTimeseriesWidget(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.TimeseriesWidget{}
	expected.X = datadog.Int(1)
	expected.Y = datadog.Int(1)
	expected.Width = datadog.Int(20)
	expected.Height = datadog.Int(30)
	expected.Title = datadog.Bool(true)
	expected.TitleAlign = datadog.String("centre")
	expected.TitleSize = &datadog.TextSize{Size: datadog.Int(16)}
	expected.TitleText = datadog.String("Test")
	expected.Timeframe = datadog.String("1m")

	w := datadog.Widget{TimeseriesWidget: expected}

	board.Widgets = append(board.Widgets, w)
	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].TimeseriesWidget

	assert.Equal(t, *actualWidget.X, *expected.X)
	assert.Equal(t, *actualWidget.Y, *expected.Y)
	assert.Equal(t, *actualWidget.Height, *expected.Height)
	assert.Equal(t, *actualWidget.Width, *expected.Width)
	assert.Equal(t, *actualWidget.TitleText, *expected.TitleText)
	assert.Equal(t, *actualWidget.TitleSize, *expected.TitleSize)
	assert.Equal(t, *actualWidget.TitleAlign, *expected.TitleAlign)
	assert.Equal(t, *actualWidget.Title, *expected.Title)

	assert.Equal(t, *actualWidget.TitleSize.Size, *expected.TitleSize.Size)
	assert.Equal(t, *actualWidget.TitleSize.Auto, *expected.TitleSize.Auto)
	assert.Equal(t, *actualWidget.Type, *expected.Type)
	assert.Equal(t, *actualWidget.Timeframe, *expected.Timeframe)
	assert.Equal(t, *actualWidget.Legend, *expected.Legend)
	assert.Equal(t, *actualWidget.TileDef, *expected.TileDef)
}

func TestQueryValueWidget(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.QueryValueWidget{}

	expected.X = datadog.Int(1)
	expected.Y = datadog.Int(1)
	expected.Width = datadog.Int(20)
	expected.Height = datadog.Int(30)
	expected.Title = datadog.Bool(true)
	expected.TitleAlign = datadog.String("centre")
	expected.TitleSize = &datadog.TextSize{Size: datadog.Int(16)}
	expected.TitleText = datadog.String("Test")
	expected.Timeframe = datadog.String("1m")

	expected.TimeframeAggregator = datadog.String("sum")
	expected.Aggregator = datadog.String("min")
	expected.Query = datadog.String("docker.containers.running")
	expected.MetricType = datadog.String("standard")
	/* TODO: add test for conditional formats
	"conditional_formats": [{
		"comparator": ">",
		"color": "white_on_red",
		"custom_bg_color": null,
		"value": 1,
		"invert": false,
		"custom_fg_color": null}],
	*/
	expected.IsValidQuery = datadog.Bool(true)
	expected.ResultCalcFunc = datadog.String("raw")
	expected.Aggregator = datadog.String("avg")
	expected.CalcFunc = datadog.String("raw")

	w := datadog.Widget{QueryValueWidget: expected}

	board.Widgets = append(board.Widgets, w)
	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := actual.Widgets[0].QueryValueWidget

	assert.Equal(t, *actualWidget.X, *expected.X)
	assert.Equal(t, *actualWidget.Y, *expected.Y)
	assert.Equal(t, *actualWidget.Height, *expected.Height)
	assert.Equal(t, *actualWidget.Width, *expected.Width)
	assert.Equal(t, *actualWidget.TitleText, *expected.TitleText)
	assert.Equal(t, *actualWidget.TitleSize, *expected.TitleSize)
	assert.Equal(t, *actualWidget.TitleAlign, *expected.TitleAlign)
	assert.Equal(t, *actualWidget.Title, *expected.Title)

	assert.Equal(t, *actualWidget.Type, *expected.Type)
	assert.Equal(t, *actualWidget.Timeframe, *expected.Timeframe)
	assert.Equal(t, *actualWidget.TimeframeAggregator, *expected.TimeframeAggregator)
	assert.Equal(t, *actualWidget.Aggregator, *expected.Aggregator)
	assert.Equal(t, *actualWidget.Query, *expected.Query)
	assert.Equal(t, *actualWidget.IsValidQuery, *expected.IsValidQuery)
	assert.Equal(t, *actualWidget.ResultCalcFunc, *expected.ResultCalcFunc)
	assert.Equal(t, *actualWidget.Aggregator, *expected.Aggregator)
}

func assertTileDefEquals(t *testing.T, actual datadog.TileDef, expected datadog.TileDef) {
	assert.Equal(t, len(actual.Events), len(expected.Events))
	assert.Equal(t, len(actual.Events), len(expected.Events))
	assert.Equal(t, len(actual.Requests), len(expected.Requests))
	if actual.Viz != nil && expected.Viz != nil {
		assert.Equal(t, *actual.Viz, *expected.Viz)
	}

	for i, event := range actual.Events {
		assert.Equal(t, *event.Query, *expected.Events[i].Query)
	}

	for i, request := range actual.Requests {
		assert.Equal(t, *request.Query, *expected.Requests[i].Query)
		assert.Equal(t, *request.Type, *expected.Requests[i].Type)
	}
}
