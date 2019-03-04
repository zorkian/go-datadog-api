package datadog

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func TestGetScreenboard(t *testing.T) {
	tests := []struct {
		file string
		want *Screenboard
	}{
		{
			file: "screenboard_response",
			want: &Screenboard{
				Id:       Int(6334),
				Title:    String("dogapi test"),
				Height:   StrInt("768"),
				Width:    StrInt("1024"),
				ReadOnly: Bool(false),
				Widgets: []Widget{
					{
						Type:   String("image"),
						Height: Int(20),
						Width:  Int(32),
						X:      Float32(32),
						Y:      Float32(7),
						URL:    String("http://path/to/image.jpg"),
					},
					{
						Type: String("timeseries"),
						TileDef: &TileDef{
							Precision: StrInt("42"),
						},
					},
					{
						Type: String("timeseries"),
						TileDef: &TileDef{
							Precision: StrInt("*"),
						},
					},
				},
			},
		},
		{
			file: "screenboard_response_fullscreen",
			want: &Screenboard{
				Id:       Int(6335),
				Title:    String("dogapi fullscreen test"),
				Height:   StrInt("100%"),
				Width:    StrInt("100%"),
				ReadOnly: Bool(false),
			},
		},
		{
			file: "screenboard_response_manage_status",
			want: &Screenboard{
				Id:    Int(6336),
				Title: String("dogapi manage_status test"),
				Widgets: []Widget{
					{
						Type: String("manage_status"),
						Params: &Params{
							Sort:  String("status,asc"),
							Text:  String(`scope:"priority:important" muted:false`),
							Count: StrInt("50"),
							Start: StrInt("0"),
						},
					},
				},
			},
		},
		{
			file: "screenboard_response_conditional_format",
			want: &Screenboard{
				Id:    Int(334488),
				Title: String("OrderCapture ScreenBoard"),
				Widgets: []Widget{
					{
						Type: String("query_value"),
						TileDef: &TileDef{
							Viz: String("query_value"),
							Requests: []TileDefRequest{
								{
									ConditionalFormats: []ConditionalFormat{
										// string values in JSON
										{
											Palette:    String("white_on_red"),
											Comparator: String(">"),
											Value:      StrInt("1000"),
										},
										{
											Palette:    String("white_on_yellow"),
											Comparator: String(">="),
											Value:      StrInt("200"),
										},
										{
											Palette:    String("white_on_green"),
											Comparator: String("<"),
											Value:      StrInt("200"),
										},
									},
								},
							},
						},
					},
					{
						Type: String("query_value"),
						TileDef: &TileDef{
							Viz: String("query_value"),
							Requests: []TileDefRequest{
								{
									ConditionalFormats: []ConditionalFormat{
										// null values in JSON
										{
											Palette:    String("white_on_red"),
											Comparator: String(">"),
											Value:      nil,
										},
										{
											Palette:    String("white_on_yellow"),
											Comparator: String(">="),
											Value:      nil,
										},
										{
											Palette:    String("white_on_green"),
											Comparator: String("<"),
											Value:      nil,
										},
									},
								},
							},
						},
					},
					{
						Type: String("toplist"),
						TileDef: &TileDef{
							Viz: String("toplist"),
							Requests: []TileDefRequest{
								{
									ConditionalFormats: []ConditionalFormat{
										// number value in JSON
										{
											Palette:    String("white_on_red"),
											Comparator: String(">"),
											Value:      StrInt("1"),
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			file: "screenboard_response_widget_float_position",
			want: &Screenboard{
				Widgets: []Widget{
					{
						Type: String("free_text"),
						Text: String("processed"),
						X:    Float32(1),
						Y:    Float32(70.83333333333334),
					},
				},
			},
		},
		{
			file: "screenboard_response_manage_status_titlesize",
			want: &Screenboard{
				Widgets: []Widget{
					{
						Type: String("manage_status"),

						ManageStatusTitleSize: StrInt("16"),
					},
				},
			},
		},
		{
			file: "screenboard_response_marker_label",
			want: &Screenboard{
				Widgets: []Widget{
					{
						Type: String("timeseries"),
						TileDef: &TileDef{
							Markers: []TileDefMarker{
								{
									Label: StrBool("true"),
								},
							},
						},
					},
				},
			},
		},
		{
			file: "screenboard_response_palette_flip",
			want: &Screenboard{
				Widgets: []Widget{
					{
						Type: String("hostmap"),
						TileDef: &TileDef{
							Viz: String("hostmap"),
							Style: &TileDefStyle{
								PaletteFlip: StrBool("false"),
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				response, err := ioutil.ReadFile("./testdata/fixtures/" + tt.file + ".json")
				if err != nil {
					// mustn't call t.Fatal from a different Goroutine
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(response)
			}))
			defer ts.Close()

			datadogClient := Client{
				baseUrl:      ts.URL,
				HttpClient:   ts.Client(),
				RetryTimeout: 5 * time.Second,
			}

			got, err := datadogClient.GetScreenboard(6334)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				// Go formatting doesn't handle pointers gracefully, so we're spewing
				t.Errorf("got:\n%s\nwant:\n%s", spew.Sdump(got), spew.Sdump(tt.want))
			}
		})
	}
}
