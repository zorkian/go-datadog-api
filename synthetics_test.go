package datadog

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchSyntheticsChecks(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/synthetics_check_search_response.json")
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

	query := ""
	checks, err := datadogClient.SearchSyntheticsChecks(query)
	if err != nil {
		t.Fatal(err)
	}

	expectedCnt := 3
	if cnt := len(checks); cnt != expectedCnt {
		t.Fatalf("expect %d checks. Got %d", expectedCnt, cnt)
	}
}
