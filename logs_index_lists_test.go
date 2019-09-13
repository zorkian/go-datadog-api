package datadog

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogsIndexListGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/logs/indexlist_response.json")
		assert.Nil(t, err)
		w.Write(response)
	}))

	defer ts.Close()

	client := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	indexList, err := client.GetLogsIndexList()

	assert.Nil(t, err)

	assert.Equal(t, []string{"main", "minor"}, indexList.IndexNames)

}
