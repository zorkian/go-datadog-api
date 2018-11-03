package datadog

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetScreenboard(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/screenboard_response.json")
		if err != nil {
			t.Fatal(err)
		}
		w.Write(response)
	}))
	defer ts.Close()

	datadogClient := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	screenboard, err := datadogClient.GetScreenboard(6334)
	if err != nil {
		t.Fatal(err)
	}

	expectedID := 6334
	if id := screenboard.GetId(); id != expectedID {
		t.Fatalf("expect ID %d. Got %d", expectedID, id)
	}

	expectedTitle := "dogapi test"
	if title := screenboard.GetTitle(); title != expectedTitle {
		t.Fatalf("expect title %s. Got %s", expectedTitle, title)
	}

	expectedHeight := 768
	if height := screenboard.GetHeight(); height != expectedHeight {
		t.Fatalf("expect height %d. Got %d", expectedHeight, height)
	}

	expectedWidth := 1024
	if width := screenboard.GetWidth(); width != expectedWidth {
		t.Fatalf("expect width %d. Got %d", expectedWidth, width)
	}

	expectedReadOnly := false
	readOnly, ok := screenboard.GetReadOnlyOk()
	if !ok {
		t.Fatalf("expect to have a read_only field")
	}

	if readOnly != expectedReadOnly {
		t.Fatalf("expect read_only %v. Got %v", expectedReadOnly, readOnly)
	}

	for _, widget := range screenboard.Widgets {
		validateWidget(t, widget)
	}
}

func validateWidget(t *testing.T, wd Widget) {
	expectedType := "image"
	if widgetType := wd.GetType(); widgetType != expectedType {
		t.Fatalf("expect type %s. Got %s", expectedType, widgetType)
	}

	expectedHeight := 20
	if height := wd.GetHeight(); height != expectedHeight {
		t.Fatalf("expect height %d. Got %d", expectedHeight, height)
	}

	expectedWidth := 32
	if width := wd.GetWidth(); width != expectedWidth {
		t.Fatalf("expect width %d. Got %d", expectedWidth, width)
	}

	expectedX := 32
	if x := wd.GetX(); x != expectedX {
		t.Fatalf("expect x %d. Got %d", expectedX, x)
	}

	expectedY := 7
	if y := wd.GetY(); y != expectedY {
		t.Fatalf("expect y %d. Got %d", expectedY, y)
	}

	expectedURL := "http://path/to/image.jpg"
	if url := wd.GetURL(); url != expectedURL {
		t.Fatalf("expect url %s. Got %s", expectedURL, url)
	}
}
