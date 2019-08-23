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
