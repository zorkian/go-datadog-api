package datadog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLogsList(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/logs/loglist_response.json")
		assert.Nil(err)
		w.Write(response)
	}))

	defer ts.Close()

	client := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	req := &LogsListRequest{
		Index: String("main"),
		Limit: Int(50),
	}

	logList, err := client.GetLogsList(req)

	assert.Nil(err)
	assert.Len(logList.Logs, 2)

}

func TestGetLogsListPages(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseBody := &LogsListRequest{}

		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, responseBody)

		if responseBody.StartAt == nil {
			response, err := ioutil.ReadFile("./tests/fixtures/logs/loglist_response.json")

			assert.Nil(err)
			w.Write(response)
		} else {
			assert.Equal(*responseBody.StartAt, "BBBBBWgN8Xwgr1vKDQAAAABBV2dOOFh3ZzZobm1mWXJFYTR0OA")

			response, err := ioutil.ReadFile("./tests/fixtures/logs/loglist_page_response.json")

			assert.Nil(err)
			w.Write(response)
		}
	}))

	defer ts.Close()

	client := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	req := &LogsListRequest{
		Index: String("main"),
		Limit: Int(50),
	}

	logs, err := client.GetLogsListPages(req, -1)

	assert.Nil(err)
	assert.Len(logs, 3)
}
