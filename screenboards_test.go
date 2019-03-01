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
						X:      Int(32),
						Y:      Int(7),
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
