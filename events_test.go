package datadog

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEvent(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/events_response.json")
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

	// Test a single event
	event, err := datadogClient.GetEvent(1377281704830403917)
	if err != nil {
		t.Fatal(err)
	}

	expectedID := 1377281704830403917
	if id := event.GetId(); id != expectedID {
		t.Fatalf("expect ID %d. Got %d", expectedID, id)
	}

	expectedTime := 1346355252
	if time := event.GetTime(); time != expectedTime {
		t.Fatalf("expect timestamp %d. Got %d", expectedTime, time)
	}

	expectedText := "Oh wow!"
	if text := event.GetText(); text != expectedText {
		t.Fatalf("expect text %s. Got %s", expectedText, text)
	}

	expectedTitle := "Did you hear the news today?"
	if title := event.GetTitle(); title != expectedTitle {
		t.Fatalf("expect title %s. Got %s", expectedTitle, title)
	}

	expectedAlertType := "info"
	if alertType := event.GetAlertType(); alertType != expectedAlertType {
		t.Fatalf("expect alert type %s. Got %s", expectedAlertType, alertType)
	}

	// host is null
	expectedHost := ""
	if host := event.GetHost(); host != expectedHost {
		t.Fatalf("expect host %s. Got %s", expectedHost, host)
	}

	expectedResource := "/api/v1/events/1377281704830403917"
	if resource := event.GetResource(); resource != expectedResource {
		t.Fatalf("expect resource %s. Got %s", expectedResource, resource)
	}

	expectedURL := "/event/jump_to?event_id=1377281704830403917"
	if url := event.GetUrl(); url != expectedURL {
		t.Fatalf("expect url %s. Got %s", expectedURL, url)
	}
}
