package datadog

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_isRateLimited(t *testing.T) {
	tests := []struct {
		desc             string
		endpoint         string
		method           string
		isRateLimited    bool
		endpointFormated string
	}{
		{
			desc:             "is rate limited",
			endpoint:         "/v1/query?&query=avg:system.cpu.user{*}by{host}",
			method:           "GET",
			isRateLimited:    true,
			endpointFormated: "/v1/query",
		},
		{
			desc:             "is not rate limited",
			endpoint:         "/v1/series?api_key=12",
			method:           "POST",
			isRateLimited:    false,
			endpointFormated: "",
		},
		{
			desc:             "is rate limited but wrong method",
			endpoint:         "/v1/query?&query=avg:system.cpu.user{*}by{host}",
			method:           "POST",
			isRateLimited:    false,
			endpointFormated: "",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d %s", i, tt.desc), func(t *testing.T) {
			limited, edpt := isRateLimited(tt.method, tt.endpoint)
			assert.Equal(t, limited, tt.isRateLimited)
			assert.Equal(t, edpt, tt.endpointFormated)
		})
	}
}

func Test_updateRateLimits(t *testing.T) {
	// fake client to ensure that we are race free.
	client := Client{
		rateLimitingStats: make(map[string]RateLimit),
	}
	tests := []struct {
		desc   string
		api    string
		resp   *http.Response
		header RateLimit
		error  error
	}{
		{
			"nominal case query",
			"/v1/query",
			makeHeader("1", "2", "3", "4"),
			RateLimit{"1", "2", "3", "4"},
			nil,
		},
		{
			"nominal case logs",
			"/v1/logs-queries/list",
			makeHeader("2", "2", "1", "5"),
			RateLimit{"2", "2", "1", "5"},
			nil,
		},
		{
			"no response",
			"",
			nil,
			RateLimit{},
			fmt.Errorf("could not parse headers from the HTTP response."),
		},
		{
			"no header",
			"/v2/error",
			makeEmptyHeader(),
			RateLimit{},
			fmt.Errorf("could not parse headers from the HTTP response."),
		},
		{
			"update case query",
			"/v1/query",
			makeHeader("2", "4", "6", "4"),
			RateLimit{"2", "4", "6", "4"},
			nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d %s", i, tt.desc), func(t *testing.T) {
			err := client.updateRateLimits(tt.resp, tt.api)
			assert.Equal(t, tt.error, err)
			assert.Equal(t, tt.header, client.rateLimitingStats[tt.api])
		})
	}
}

func makeHeader(limit, period, reset, remaining string) *http.Response {
	h := http.Response{
		Header: make(map[string][]string),
	}
	h.Header.Set("X-RateLimit-Limit", limit)
	h.Header.Set("X-RateLimit-Reset", reset)
	h.Header.Set("X-RateLimit-Period", period)
	h.Header.Set("X-RateLimit-Remaining", remaining)
	return &h
}

func makeEmptyHeader() *http.Response {
	return &http.Response{
		Header: nil,
	}
}
