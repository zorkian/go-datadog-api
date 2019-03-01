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
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./testdata/fixtures/screenboard_response.json")
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

	want := &Screenboard{
		Id:       Int(6334),
		Title:    String("dogapi test"),
		Height:   Int(768),
		Width:    Int(1024),
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
		},
	}

	if !reflect.DeepEqual(got, want) {
		// Go formatting doesn't handle pointers gracefully, so we're spewing
		t.Errorf("got:\n%s\nwant:\n%s", spew.Sdump(got), spew.Sdump(want))
	}
}
