package datadog

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHostTotals(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/hosts/get_totals_response.json")
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

	res, err := datadogClient.GetHostTotals()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, *res.TotalActive, 1)
	assert.Equal(t, *res.TotalUp, 2)
}

func TestGetHosts(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/hosts/get_hosts_response.json")
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

	searchRequest := HostSearchRequest{
		SortField:     "cpu",
		SortDirection: "desc",
		Start:         0,
		Count:         100,
	}
	res, err := datadogClient.GetHosts(searchRequest)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.TotalReturned, 3)
	assert.Equal(t, res.TotalMatching, 3)
	assert.Len(t, res.HostList, 3)
}
