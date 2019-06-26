package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
)

func TestWidgets(t *testing.T) {
	widgets := []datadog.Widget{
		{
			Type:      datadog.String("free_text"),
			X:         datadog.Int(1),
			Y:         datadog.Int(1),
			Width:     datadog.Int(5),
			Height:    datadog.Int(5),
			Text:      datadog.String("Test"),
			TextAlign: datadog.String("right"),
			FontSize:  datadog.String("36"),
			Color:     datadog.String("#ffc0cb"),
		},
		{
			Type:       datadog.String("timeseries"),
			X:          datadog.Int(1),
			Y:          datadog.Int(1),
			Width:      datadog.Int(5),
			Height:     datadog.Int(5),
			Title:      datadog.Bool(true),
			TitleText:  datadog.String("Test title"),
			TitleSize:  datadog.Int(16),
			TitleAlign: datadog.String("right"),
			Legend:     datadog.Bool(true),
			LegendSize: datadog.String("16"),
			Time: &datadog.Time{
				LiveSpan: datadog.String("1d"),
			},
			TileDef: &datadog.TileDef{
				Viz: datadog.String("timeseries"),
				Requests: []datadog.TileDefRequest{{
					Query: datadog.String("avg:system.cpu.user{*}"),
					Type:  datadog.String("line"),
					Style: &datadog.TileDefRequestStyle{
						Palette: datadog.String("purple"),
						Type:    datadog.String("dashed"),
						Width:   datadog.String("thin"),
					},
					Metadata: map[string]datadog.TileDefMetadata{
						"avg:system.cpu.user{*}": {
							Alias: datadog.String("avg_cpu"),
						},
					},
				}},
				Markers: []datadog.TileDefMarker{{
					Label: datadog.String("test marker"),
					Type:  datadog.String("error dashed"),
					Value: datadog.String("y < 6"),
				}},
				Events: []datadog.TileDefEvent{{
					Query: datadog.String("test event"),
				}},
			},
		},
		{
			Type:       datadog.String("query_value"),
			X:          datadog.Int(1),
			Y:          datadog.Int(1),
			Width:      datadog.Int(5),
			Height:     datadog.Int(5),
			Title:      datadog.Bool(true),
			TitleText:  datadog.String("Test title"),
			TitleSize:  datadog.Int(16),
			TitleAlign: datadog.String("right"),
			Legend:     datadog.Bool(true),
			LegendSize: datadog.String("16"),
			Time: &datadog.Time{
				LiveSpan: datadog.String("1d"),
			},
			TileDef: &datadog.TileDef{
				Viz: datadog.String("query_value"),
				Requests: []datadog.TileDefRequest{{
					Query: datadog.String("avg:system.cpu.user{*}"),
					Type:  datadog.String("line"),
					Style: &datadog.TileDefRequestStyle{
						Palette: datadog.String("purple"),
						Type:    datadog.String("dashed"),
						Width:   datadog.String("thin"),
					},
					ConditionalFormats: []datadog.ConditionalFormat{
						{
							Comparator:    datadog.String(">="),
							CustomBgColor: datadog.String("#205081"),
							Value:         datadog.String("1"),
							Palette:       datadog.String("white_on_red"),
						}},
					Aggregator: datadog.String("max"),
				}},
				CustomUnit: datadog.String("%"),
				Autoscale:  datadog.Bool(false),
				Precision:  datadog.Precision("6"),
				TextAlign:  datadog.String("right"),
			},
		},
		{
			Type:       datadog.String("toplist"),
			X:          datadog.Int(1),
			Y:          datadog.Int(1),
			Width:      datadog.Int(5),
			Height:     datadog.Int(5),
			Title:      datadog.Bool(true),
			TitleText:  datadog.String("Test title"),
			TitleSize:  datadog.Int(16),
			TitleAlign: datadog.String("right"),
			Legend:     datadog.Bool(true),
			LegendSize: datadog.String("16"),
			Time: &datadog.Time{
				LiveSpan: datadog.String("1d"),
			},
			TileDef: &datadog.TileDef{
				Viz: datadog.String("toplist"),
				Requests: []datadog.TileDefRequest{{
					Query: datadog.String("top(avg:system.load.1{*} by {host}, 10, 'mean', 'desc')"),
					Style: &datadog.TileDefRequestStyle{
						Palette: datadog.String("purple"),
						Type:    datadog.String("dashed"),
						Width:   datadog.String("thin"),
					},
					ConditionalFormats: []datadog.ConditionalFormat{
						{
							Comparator: datadog.String(">"),
							Value:      datadog.String("4"),
							Palette:    datadog.String("white_on_green"),
						}},
				}},
			},
		},
		{
			Type:       datadog.String("change"),
			X:          datadog.Int(1),
			Y:          datadog.Int(1),
			Width:      datadog.Int(5),
			Height:     datadog.Int(5),
			Title:      datadog.Bool(true),
			TitleText:  datadog.String("Test title"),
			TitleSize:  datadog.Int(16),
			TitleAlign: datadog.String("right"),
			Time: &datadog.Time{
				LiveSpan: datadog.String("1d"),
			},
			TileDef: &datadog.TileDef{
				Viz: datadog.String("change"),
				Requests: []datadog.TileDefRequest{{
					Query:        datadog.String("min:system.load.1{*} by {host}"),
					CompareTo:    datadog.String("week_before"),
					ChangeType:   datadog.String("relative"),
					OrderBy:      datadog.String("present"),
					OrderDir:     datadog.String("asc"),
					ExtraCol:     datadog.String(""),
					IncreaseGood: datadog.Bool(false),
				}},
			},
		},
		{
			Type:       datadog.String("event_timeline"),
			X:          datadog.Int(1),
			Y:          datadog.Int(1),
			Width:      datadog.Int(5),
			Height:     datadog.Int(5),
			Title:      datadog.Bool(true),
			TitleText:  datadog.String("Test title"),
			TitleSize:  datadog.Int(16),
			TitleAlign: datadog.String("right"),
			Time: &datadog.Time{
				LiveSpan: datadog.String("1d"),
			},
		},
		{
			Type:       datadog.String("event_stream"),
			X:          datadog.Int(1),
			Y:          datadog.Int(1),
			Width:      datadog.Int(5),
			Height:     datadog.Int(5),
			Title:      datadog.Bool(true),
			TitleText:  datadog.String("Test title"),
			TitleSize:  datadog.Int(16),
			TitleAlign: datadog.String("right"),
			Query:      datadog.String("*"),
			EventSize:  datadog.String("l"),
			Time: &datadog.Time{
				LiveSpan: datadog.String("4h"),
			},
		},
		{
			Type:       datadog.String("image"),
			X:          datadog.Int(1),
			Y:          datadog.Int(1),
			Width:      datadog.Int(5),
			Height:     datadog.Int(5),
			Title:      datadog.Bool(true),
			TitleText:  datadog.String("Test title"),
			TitleSize:  datadog.Int(16),
			TitleAlign: datadog.String("right"),
			Sizing:     datadog.String("fit"),
			Margin:     datadog.String("large"),
			URL:        datadog.String("https://datadog-prod.imgix.net/img/dd_logo_70x75.png"),
		},
		{
			Type:      datadog.String("note"),
			X:         datadog.Int(1),
			Y:         datadog.Int(1),
			Width:     datadog.Int(5),
			Height:    datadog.Int(5),
			Bgcolor:   datadog.String("pink"),
			TextAlign: datadog.String("right"),
			FontSize:  datadog.String("36"),
			Tick:      datadog.Bool(true),
			TickEdge:  datadog.String("bottom"),
			TickPos:   datadog.String("50%"),
			HTML:      datadog.String("<b>test</b>"),
		},
		{
			Type:       datadog.String("alert_graph"),
			X:          datadog.Int(1),
			Y:          datadog.Int(1),
			Width:      datadog.Int(5),
			Height:     datadog.Int(5),
			Title:      datadog.Bool(true),
			TitleText:  datadog.String("Test title"),
			TitleSize:  datadog.Int(16),
			TitleAlign: datadog.String("right"),
			AlertID:    datadog.Int(123456),
			VizType:    datadog.String("toplist"),
			Time: &datadog.Time{
				LiveSpan: datadog.String("15m"),
			},
		},
		{
			Type:       datadog.String("alert_value"),
			X:          datadog.Int(1),
			Y:          datadog.Int(1),
			Width:      datadog.Int(5),
			Height:     datadog.Int(5),
			Title:      datadog.Bool(true),
			TitleText:  datadog.String("Test title"),
			TitleSize:  datadog.Int(16),
			TitleAlign: datadog.String("right"),
			AlertID:    datadog.Int(123456),
			TextSize:   datadog.String("fill_height"),
			TextAlign:  datadog.String("right"),
			Precision:  datadog.Precision("*"),
			Unit:       datadog.String("b"),
		},
		{
			Type:   datadog.String("iframe"),
			X:      datadog.Int(1),
			Y:      datadog.Int(1),
			Width:  datadog.Int(5),
			Height: datadog.Int(5),
			URL:    datadog.String("https://www.datadoghq.com/"),
		},
		{
			Type:       datadog.String("check_status"),
			X:          datadog.Int(1),
			Y:          datadog.Int(1),
			Width:      datadog.Int(5),
			Height:     datadog.Int(5),
			Title:      datadog.Bool(true),
			TitleText:  datadog.String("Test title"),
			TitleSize:  datadog.Int(16),
			TitleAlign: datadog.String("right"),
			Grouping:   datadog.String("check"),
			Check:      datadog.String("aws.ecs.agent_connected"),
			Tags:       []*string{datadog.String("*")},
			Group:      datadog.String("cluster:test"),
			Time: &datadog.Time{
				LiveSpan: datadog.String("15m"),
			},
		},
		{
			Type:                 datadog.String("trace_service"),
			X:                    datadog.Int(1),
			Y:                    datadog.Int(1),
			Width:                datadog.Int(5),
			Height:               datadog.Int(5),
			Env:                  datadog.String("test"),
			ServiceService:       datadog.String("service"),
			ServiceName:          datadog.String("serviceName"),
			SizeVersion:          datadog.String("large"),
			LayoutVersion:        datadog.String("three_column"),
			MustShowHits:         datadog.Bool(true),
			MustShowErrors:       datadog.Bool(true),
			MustShowLatency:      datadog.Bool(true),
			MustShowBreakdown:    datadog.Bool(true),
			MustShowDistribution: datadog.Bool(true),
			MustShowResourceList: datadog.Bool(true),
			Time: &datadog.Time{
				LiveSpan: datadog.String("15m"),
			},
		},
		{
			Type:   datadog.String("hostmap"),
			X:      datadog.Int(1),
			Y:      datadog.Int(1),
			Width:  datadog.Int(5),
			Height: datadog.Int(5),
			Query:  datadog.String("avg:system.load.1{*} by {host}"),
			TileDef: &datadog.TileDef{
				Viz:           datadog.String("hostmap"),
				NodeType:      datadog.String("container"),
				Scope:         []*string{datadog.String("tag:test")},
				Group:         []*string{datadog.String("test")},
				NoGroupHosts:  datadog.Bool(false),
				NoMetricHosts: datadog.Bool(false),
				Requests: []datadog.TileDefRequest{{
					Query: datadog.String("min:process.stat.container.io.wbps{tag:test} by {host}"),
					Type:  datadog.String("fill"),
				}},
				Style: &datadog.TileDefStyle{
					Palette:     datadog.String("hostmap_blues"),
					PaletteFlip: datadog.String("true"),
					FillMin:     datadog.JsonNumber("20"),
					FillMax:     datadog.JsonNumber("300"),
				},
			},
		},
		{
			Type:                   datadog.String("manage_status"),
			X:                      datadog.Int(1),
			Y:                      datadog.Int(1),
			Width:                  datadog.Int(5),
			Height:                 datadog.Int(5),
			DisplayFormat:          datadog.String("countsAndList"),
			ColorPreference:        datadog.String("background"),
			HideZeroCounts:         datadog.Bool(true),
			ManageStatusShowTitle:  datadog.Bool(false),
			ManageStatusTitleText:  datadog.String("Test title"),
			ManageStatusTitleSize:  datadog.String("20"),
			ManageStatusTitleAlign: datadog.String("right"),
			Params: &datadog.Params{
				Sort:  datadog.String("status,asc"),
				Text:  datadog.String("status:alert"),
				Count: datadog.String("50"),
				Start: datadog.String("0"),
			},
		},
		{
			Type:    datadog.String("log_stream"),
			X:       datadog.Int(1),
			Y:       datadog.Int(1),
			Width:   datadog.Int(5),
			Height:  datadog.Int(5),
			Query:   datadog.String("source:main"),
			Columns: datadog.String("[\"column_1\",\"column_2\",\"column_3\"]"),
			Logset:  datadog.String("1234"),
			Time: &datadog.Time{
				LiveSpan: datadog.String("1h"),
			},
		},
		{
			// Widget is undocumented, subject to breaking API changes, and without customer support
			Type:   datadog.String("uptime"),
			X:      datadog.Int(1),
			Y:      datadog.Int(1),
			Width:  datadog.Int(5),
			Height: datadog.Int(5),
			Timeframes: []*string{
				datadog.String("7 days"),
				datadog.String("Month-to-date"),
				datadog.String("90 days"),
			},
			Rules: map[string]*datadog.Rule{
				"0": {
					Threshold: datadog.JsonNumber("95"),
					Timeframe: datadog.String("Month-to-date"),
					Color:     datadog.String("green"),
				},
				"1": {
					Threshold: datadog.JsonNumber("98"),
					Timeframe: datadog.String("7 days"),
					Color:     datadog.String("red"),
				},
			},
			Monitor: &datadog.ScreenboardMonitor{
				Id: datadog.Int(1234),
			},
		},
		{
			Type:   datadog.String("process"),
			X:      datadog.Int(1),
			Y:      datadog.Int(1),
			Width:  datadog.Int(5),
			Height: datadog.Int(5),
			Time:   &datadog.Time{},
			TileDef: &datadog.TileDef{
				Viz: datadog.String("process"),
				Requests: []datadog.TileDefRequest{{
					QueryType:  datadog.String("process"),
					Metric:     datadog.String("process.stat.cpu.total_pct"),
					TextFilter: datadog.String("test"),
					TagFilters: []*string{datadog.String("test")},
					Limit:      datadog.Int(200),
					Style: &datadog.TileDefRequestStyle{
						Palette: datadog.String("dog_classic_area"),
					},
				}},
			},
		},
	}

	for _, w := range widgets {
		t.Run(*w.Type, func(t *testing.T) {
			board := createTestScreenboard(t)
			defer cleanUpScreenboard(t, *board.Id)

			board.Widgets = append(board.Widgets, w)

			if err := client.UpdateScreenboard(board); err != nil {
				t.Fatalf("Updating a screenboard failed: %s", err)
			}

			actual, err := client.GetScreenboard(*board.Id)
			if err != nil {
				t.Fatalf("Retrieving a screenboard failed: %s", err)
			}

			actualWidget := actual.Widgets[0]

			assert.Equal(t, actualWidget, w)
		})
	}
}
