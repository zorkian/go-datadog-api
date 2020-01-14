package datadog

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func Test_updateRateLimits(t *testing.T) {
	// fake client to ensure that we are race free.
	client := Client{
		rateLimitingStats: make(map[string]RateLimit),
	}
	tests := []struct {
		desc   string
		api    *url.URL
		resp   *http.Response
		header RateLimit
		error  error
	}{
		{
			"nominal case query",
			&url.URL{Path: "/v1/query"},
			makeHeader(RateLimit{"1", "2", "3", "4"}),
			RateLimit{"1", "2", "3", "4"},
			nil,
		},
		{
			"nominal case logs",
			&url.URL{Path: "/v1/logs-queries/list"},
			makeHeader(RateLimit{"2", "2", "1", "5"}),
			RateLimit{"2", "2", "1", "5"},
			nil,
		},
		{
			"no response",
			&url.URL{Path: ""},
			nil,
			RateLimit{},
			fmt.Errorf("malformed HTTP content."),
		},
		{
			"no header",
			&url.URL{Path: "/v2/error"},
			makeEmptyHeader(),
			RateLimit{},
			fmt.Errorf("malformed HTTP content."),
		},
		{
			"not rate limited",
			&url.URL{Path: "/v2/error"},
			makeHeader(RateLimit{}),
			RateLimit{},
			nil,
		},
		{
			"update case query",
			&url.URL{Path: "/v1/query"},
			makeHeader(RateLimit{"2", "4", "6", "4"}),
			RateLimit{"2", "4", "6", "4"},
			nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d %s", i, tt.desc), func(t *testing.T) {
			err := client.updateRateLimits(tt.resp, tt.api)
			assert.Equal(t, tt.error, err)
			assert.Equal(t, tt.header, client.rateLimitingStats[tt.api.Path])
		})
	}
}

func makeHeader(r RateLimit) *http.Response {
	h := http.Response{
		Header: make(map[string][]string),
	}
	h.Header.Set("X-RateLimit-Limit", r.Limit)
	h.Header.Set("X-RateLimit-Reset", r.Reset)
	h.Header.Set("X-RateLimit-Period", r.Period)
	h.Header.Set("X-RateLimit-Remaining", r.Remaining)
	return &h
}

func makeEmptyHeader() *http.Response {
	return &http.Response{
		Header: nil,
	}
}
