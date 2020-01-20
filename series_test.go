package datadog

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeWhitespace(s string) string {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	return s
}

// TestPostMetrics tests submitting series sends correct
// payloads to the Datadog API for the /v1/series endpoint
func TestPostMetrics(t *testing.T) {
	reqs := make(chan string, 1)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		reqs <- buf.String()
		w.WriteHeader(200)
		w.Write([]byte("{\"status\": \"ok\"}"))
		return
	}))
	defer ts.Close()

	client := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	tcs := []string{
		"./tests/fixtures/series/post_series_mixed.json",
		"./tests/fixtures/series/post_series_valid.json",
	}

	for _, tc := range tcs {
		b, err := ioutil.ReadFile(tc)
		if err != nil {
			t.Fatal(err)
		}

		var post reqPostSeries
		json.Unmarshal(b, &post)
		if err != nil {
			t.Fatal(err)
		}

		err = client.PostMetrics(post.Series)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, nil, err)

		payload := <-reqs
		assert.Equal(t, removeWhitespace(string(b)), payload)
	}

	// Empty slice metrics test case

	err := client.PostMetrics([]Metric{})
	if err != nil {
		t.Fatal(err)
	}

	payload := <-reqs
	assert.Equal(t, "{}", payload)
}
