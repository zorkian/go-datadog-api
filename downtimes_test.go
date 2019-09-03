package datadog

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDowntime(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/downtimes_response.json")
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

	downtime, err := datadogClient.GetDowntime(2910)
	if err != nil {
		t.Fatal(err)
	}

	expectedID := 2910
	if id := downtime.GetId(); id != expectedID {
		t.Fatalf("expect ID %d. Got %d", expectedID, id)
	}

	expectedActive := true
	if active := downtime.GetActive(); active != expectedActive {
		t.Fatalf("expect active %t. Got %v", expectedActive, active)
	}

	expectedEnd := 1420447087
	if end := downtime.GetEnd(); end != expectedEnd {
		t.Fatalf("expect end %d. Got %d", expectedEnd, end)
	}

	expectedDisabled := false
	if disabled := downtime.GetDisabled(); expectedDisabled != disabled {
		t.Fatalf("expect active %t. Got %v", expectedActive, disabled)
	}

	expectedMessage := "Doing some testing on staging."
	if message := downtime.GetMessage(); expectedMessage != message {
		t.Fatalf("expect message %s. Got %s", expectedMessage, message)
	}

	expectedStart := 1420387032
	if start := downtime.GetStart(); expectedStart != start {
		t.Fatalf("expect end %d. Got %d", expectedStart, start)
	}
}
