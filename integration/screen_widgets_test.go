package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
	"github.com/mitchellh/mapstructure"
)

func TestWidgetAlertValue(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.AlertValueWidget{
		X:            datadog.Int(1),
		Y:            datadog.Int(1),
		Width:        datadog.Int(5),
		Height:       datadog.Int(5),
		TitleText:    datadog.String("foo"),
		TitleAlign:   datadog.String("center"),
		TitleSize:    datadog.Int(1),
		Title:        datadog.Bool(true),
		TextSize:     datadog.String("auto"),
		Precision:    datadog.Int(2),
		AlertId:      datadog.Int(1),
		Type:         datadog.String("alert_value"),
		Unit:         datadog.String("auto"),
		AddTimeframe: datadog.Bool(false),
	}

	board.Widgets = append(board.Widgets, expected)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := datadog.AlertValueWidget{}
	mapstructure.Decode(actual.Widgets[0], &actualWidget)

	assert.Equal(t, *expected, actualWidget)

}

func TestWidgetChange(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.ChangeWidget{
		X:          datadog.Int(1),
		Y:          datadog.Int(1),
		Width:      datadog.Int(5),
		Height:     datadog.Int(5),
		TitleText:  datadog.String("foo"),
		TitleAlign: datadog.String("center"),
		TitleSize:  datadog.Int(1),
		Title:      datadog.Bool(true),
		Aggregator: datadog.String("min"),
		TileDef:    &datadog.TileDef{},
	}

	board.Widgets = append(board.Widgets, expected)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := datadog.ChangeWidget{}
	mapstructure.Decode(actual.Widgets[0], &actualWidget)

	assert.Equal(t, *expected, actualWidget)
}

func TestWidgetGraph(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.GraphWidget{
		X:          datadog.Int(1),
		Y:          datadog.Int(1),
		Width:      datadog.Int(5),
		Height:     datadog.Int(5),
		TitleText:  datadog.String("foo"),
		TitleAlign: datadog.String("center"),
		TitleSize:  datadog.Int(1),
		Title:      datadog.Bool(true),
		Time:  		&datadog.Time{
			LiveSpan: datadog.String("1d"),
		},
		Type:       datadog.String("alert_graph"),
		Legend:     datadog.Bool(true),
		LegendSize: datadog.Int(5),
		TileDef:    &datadog.TileDef{},
	}

	board.Widgets = append(board.Widgets, expected)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := datadog.GraphWidget{}
	mapstructure.Decode(actual.Widgets[0], &actualWidget)

	assert.Equal(t, *expected, actualWidget)
}

func TestWidgetEventTimeline(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.EventTimelineWidget{
		X:          datadog.Int(1),
		Y:          datadog.Int(1),
		Width:      datadog.Int(5),
		Height:     datadog.Int(5),
		TitleText:  datadog.String("foo"),
		TitleAlign: datadog.String("center"),
		TitleSize:  datadog.Int(1),
		Title:      datadog.Bool(true),
		Query:      datadog.String("avg:system.load.1{foo} by {bar}"),
		Time:  		&datadog.Time{
			LiveSpan: datadog.String("1d"),
		},
		Type:       datadog.String("alert_graph"),
	}

	board.Widgets = append(board.Widgets, expected)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := datadog.EventTimelineWidget{}
	mapstructure.Decode(actual.Widgets[0], &actualWidget)

	assert.Equal(t, *expected, actualWidget)
}

func TestAlertWidgetGraph(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.AlertGraphWidget{
		X:            datadog.Int(1),
		Y:            datadog.Int(1),
		Width:        datadog.Int(5),
		Height:       datadog.Int(5),
		TitleText:    datadog.String("foo"),
		TitleAlign:   datadog.String("center"),
		TitleSize:    datadog.Int(1),
		Title:        datadog.Bool(true),
		VizType:      datadog.String(""),
		Time:  		&datadog.Time{
			LiveSpan: datadog.String("1d"),
		},
		AddTimeframe: datadog.Bool(false),
		AlertId:      datadog.Int(1),
		Type:         datadog.String("alert_graph"),
	}

	board.Widgets = append(board.Widgets, expected)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := datadog.AlertGraphWidget{}
	mapstructure.Decode(actual.Widgets[0], &actualWidget)

	assert.Equal(t, *expected, actualWidget)
}

func TestWidgetHostMap(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.HostMapWidget{
		X:          datadog.Int(1),
		Y:          datadog.Int(1),
		Width:      datadog.Int(5),
		Height:     datadog.Int(5),
		TitleText:  datadog.String("foo"),
		TitleAlign: datadog.String("center"),
		TitleSize:  datadog.Int(1),
		Title:      datadog.Bool(true),
		Type:       datadog.String("check_status"),
		Query:      datadog.String("avg:system.load.1{foo} by {bar}"),
		Time:  		&datadog.Time{
			LiveSpan: datadog.String("1d"),
		},
		Legend:     datadog.Bool(true),
		LegendSize: datadog.Int(5),
		TileDef:    &datadog.TileDef{},
	}

	board.Widgets = append(board.Widgets, expected)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := datadog.HostMapWidget{}
	mapstructure.Decode(actual.Widgets[0], &actualWidget)

	assert.Equal(t, *expected, actualWidget)
}

func TestWidgetCheckStatus(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.CheckStatusWidget{
		X:          datadog.Int(1),
		Y:          datadog.Int(1),
		Width:      datadog.Int(5),
		Height:     datadog.Int(5),
		TitleText:  datadog.String("foo"),
		TitleAlign: datadog.String("center"),
		TitleSize:  datadog.Int(1),
		Title:      datadog.Bool(true),
		Type:       datadog.String("check_status"),
		Tags:       datadog.String("foo"),
		Time:  		&datadog.Time{
			LiveSpan: datadog.String("1d"),
		},
		Check:      datadog.String("datadog.agent.up"),
		Group:      datadog.String("foo"),
		Grouping:   datadog.String("check"),
	}

	board.Widgets = append(board.Widgets, expected)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := datadog.CheckStatusWidget{}
	mapstructure.Decode(actual.Widgets[0], &actualWidget)

	assert.Equal(t, *expected, actualWidget)
}

func TestWidgetIFrame(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.IFrameWidget{
		X:          datadog.Int(1),
		Y:          datadog.Int(1),
		Width:      datadog.Int(5),
		Height:     datadog.Int(5),
		TitleText:  datadog.String("foo"),
		TitleAlign: datadog.String("center"),
		TitleSize:  datadog.Int(1),
		Title:      datadog.Bool(true),
		Url:        datadog.String("http://www.example.com"),
		Type:       datadog.String("iframe"),
	}

	board.Widgets = append(board.Widgets, expected)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := datadog.IFrameWidget{}
	mapstructure.Decode(actual.Widgets[0], &actualWidget)

	assert.Equal(t, *expected, actualWidget)
}

func TestWidgetNote(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.NoteWidget{
		X:            datadog.Int(1),
		Y:            datadog.Int(1),
		Width:        datadog.Int(5),
		Height:       datadog.Int(5),
		TitleText:    datadog.String("foo"),
		TitleAlign:   datadog.String("center"),
		TitleSize:    datadog.Int(1),
		Title:        datadog.Bool(true),
		Color:        datadog.String("green"),
		FontSize:     datadog.Int(5),
		RefreshEvery: datadog.Int(60),
		TickPos:      datadog.String("foo"),
		TickEdge:     datadog.String("bar"),
		Html:         datadog.String("<strong>baz</strong>"),
		Tick:         datadog.Bool(false),
		Note:         datadog.String("quz"),
		AutoRefresh:  datadog.Bool(false),
	}

	board.Widgets = append(board.Widgets, expected)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := datadog.NoteWidget{}
	mapstructure.Decode(actual.Widgets[0], &actualWidget)

	assert.Equal(t, *expected, actualWidget)
}

func TestWidgetToplist(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.ToplistWidget{
		X:          datadog.Int(1),
		Y:          datadog.Int(1),
		Width:      datadog.Int(5),
		Height:     datadog.Int(5),
		TitleText:  datadog.String("foo"),
		TitleAlign: datadog.String("center"),
		TitleSize: &datadog.TextSize{
			Size: datadog.Int(1),
			Auto: datadog.Bool(true),
		},
		Title:      datadog.Bool(true),
		Time:  		&datadog.Time{
			LiveSpan: datadog.String("1d"),
		},
		Type: datadog.String("toplist"),
		Legend:     datadog.Bool(false),
		LegendSize: datadog.Int(5),
	}

	board.Widgets = append(board.Widgets, expected)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := datadog.ToplistWidget{}
	mapstructure.Decode(actual.Widgets[0], &actualWidget)

	assert.Equal(t, *expected, actualWidget)
}

func TestWidgetEventSteam(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.EventStreamWidget{
		EventSize:  datadog.String("1"),
		X:          datadog.Int(1),
		Y:          datadog.Int(1),
		Width:      datadog.Int(5),
		Height:     datadog.Int(5),
		TitleText:  datadog.String("foo"),
		TitleAlign: datadog.String("center"),
		TitleSize: &datadog.TextSize{
			Size: datadog.Int(1),
			Auto: datadog.Bool(true),
		},
		Title:     datadog.Bool(true),
		Time:  		&datadog.Time{
			LiveSpan: datadog.String("5m"),
		},
		Query:     datadog.String("tags:release"),
		Type:      datadog.String("event_stream"),
	}

	board.Widgets = append(board.Widgets, expected)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := datadog.EventStreamWidget{}
	mapstructure.Decode(actual.Widgets[0], &actualWidget)

	assert.Equal(t, *expected, actualWidget)
}

func TestWidgetImage(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.ImageWidget{
		X:          datadog.Int(1),
		Y:          datadog.Int(1),
		Width:      datadog.Int(5),
		Height:     datadog.Int(5),
		Title:      datadog.Bool(false),
		TitleAlign: datadog.String("center"),
		TitleSize: &datadog.TextSize{
			Size: datadog.Int(1),
			Auto: datadog.Bool(true),
		},
		TitleText: datadog.String("bar"),
		Type:      datadog.String("baz"),
		Url:       datadog.String("qux"),
		Sizing:    datadog.String("quuz"),
	}

	board.Widgets = append(board.Widgets, expected)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := datadog.ImageWidget{}
	mapstructure.Decode(actual.Widgets[0], &actualWidget)

	assert.Equal(t, *expected, actualWidget)
}

func TestWidgetFreeText(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.FreeTextWidget{
		X:         datadog.Int(1),
		Y:         datadog.Int(1),
		Width:     datadog.Int(5),
		Height:    datadog.Int(5),
		Text:      datadog.String("Test"),
		FontSize:  datadog.String("16"),
		TextAlign: datadog.String("center"),
		Type:      datadog.String("baz"),
	}

	board.Widgets = append(board.Widgets, expected)

	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := datadog.FreeTextWidget{}
	mapstructure.Decode(actual.Widgets[0], &actualWidget)

	assert.Equal(t, *expected, actualWidget)
}

func TestWidgetTimeseries(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.TimeseriesWidget{
		X:          datadog.Int(1),
		Y:          datadog.Int(1),
		Width:      datadog.Int(20),
		Height:     datadog.Int(30),
		Title:      datadog.Bool(true),
		TitleAlign: datadog.String("centre"),
		TitleSize: &datadog.TextSize{
			Size: datadog.Int(16),
			Auto: datadog.Bool(true),
		},
		TitleText: datadog.String("Test"),
		Type:      datadog.String("timeseries"),
		Time:  		&datadog.Time{
			LiveSpan: datadog.String("1m"),
		},
	}

	board.Widgets = append(board.Widgets, expected)
	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := datadog.TimeseriesWidget{}
	mapstructure.Decode(actual.Widgets[0], &actualWidget)

	assert.Equal(t, *expected, actualWidget)
}

func TestWidgetQueryValue(t *testing.T) {
	board := createTestScreenboard(t)
	defer cleanUpScreenboard(t, *board.Id)

	expected := &datadog.QueryValueWidget{
		X:                   datadog.Int(1),
		Y:                   datadog.Int(1),
		Width:               datadog.Int(20),
		Height:              datadog.Int(30),
		Title:               datadog.Bool(true),
		TitleAlign:          datadog.String("centre"),
		TitleSize:           &datadog.TextSize{Size: datadog.Int(16)},
		TitleText:           datadog.String("Test"),
		Time:  		&datadog.Time{
			LiveSpan: datadog.String("1m"),
		},
		TimeframeAggregator: datadog.String("sum"),
		Aggregator:          datadog.String("min"),
		Query:               datadog.String("docker.containers.running"),
		MetricType:          datadog.String("standard"),
		/* TODO: add test for conditional formats
		"conditional_formats": [{
			"comparator": ">",
			"color": "white_on_red",
			"custom_bg_color": null,
			"value": 1,
			"invert": false,
			"custom_fg_color": null}],
		*/
		IsValidQuery:   datadog.Bool(true),
		ResultCalcFunc: datadog.String("raw"),
		CalcFunc:       datadog.String("raw"),
		Type:           datadog.String("query_value"),
	}

	board.Widgets = append(board.Widgets, expected)
	if err := client.UpdateScreenboard(board); err != nil {
		t.Fatalf("Updating a screenboard failed: %s", err)
	}

	actual, err := client.GetScreenboard(*board.Id)
	if err != nil {
		t.Fatalf("Retrieving a screenboard failed: %s", err)
	}

	actualWidget := datadog.QueryValueWidget{}
	mapstructure.Decode(actual.Widgets[0], &actualWidget)

	assert.Equal(t, *expected, actualWidget)
}
