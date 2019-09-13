package datadog

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogsPipelineListGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/logs/pipelinelist_response.json")
		assert.Nil(t, err)
		w.Write(response)
	}))

	defer ts.Close()

	client := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	pipelineList, err := client.GetLogsPipelineList()

	assert.Nil(t, err)

	assert.Equal(t, 3, len(pipelineList.PipelineIds))

}
