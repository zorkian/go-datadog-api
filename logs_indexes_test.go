package datadog

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogsIndexGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/logs/index_response.json")
		if err != nil {
			t.Fatal(err)
		}
		w.Write(response)
	}))

	defer ts.Close()

	client := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	logsIndex, err := client.GetLogsIndex("main")
	assert.Nil(t, err)
	assert.Equal(t, expectedIndex, logsIndex)
}

var expectedIndex = &LogsIndex{
	Name:             String("main"),
	NumRetentionDays: Int64(90),
	DailyLimit:       Int64(151000000),
	IsRateLimited:    Bool(false),
	Filter:           &FilterConfiguration{Query: String("*")},
	ExclusionFilters: []ExclusionFilter{
		{
			Name:      String("Filter 1"),
			IsEnabled: Bool(true),
			Filter: &Filter{
				Query:      String("source:agent status:info"),
				SampleRate: Float64(1.0),
			},
		}, {
			Name:      String("Filter 2"),
			IsEnabled: Bool(true),
			Filter: &Filter{
				Query:      String("source:agent"),
				SampleRate: Float64(0.8),
			},
		}, {
			Name:      String("Filter 3"),
			IsEnabled: Bool(true),
			Filter: &Filter{
				Query:      String("source:debug"),
				SampleRate: Float64(0.5),
			},
		},
	},
}
