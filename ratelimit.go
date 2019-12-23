package datadog

import (
	"fmt"
	"net/http"
	"strings"
)

// The list of Rate Limited Endpoints of the Datadog API.
// https://docs.datadoghq.com/api/?lang=bash#rate-limiting
var (
	rateLimitedEndpoints = map[string]string{
		"/v1/query":               "GET",
		"/v1/input":               "GET",
		"/v1/metrics":             "GET",
		"/v1/events":              "POST",
		"/v1/logs-queries/list":   "POST",
		"/v1/graph/snapshot":      "GET",
		"/v1/logs/config/indexes": "GET",
	}
)

func isRateLimited(method string, endpoint string) (limited bool, shortEndpoint string) {
	for e, m := range rateLimitedEndpoints {
		if strings.HasPrefix(endpoint, e) && m == method {
			return true, e
		}
	}
	return false, ""
}

func (client *Client) updateRateLimits(resp *http.Response, api string) error {
	if resp == nil || resp.Header == nil {
		return fmt.Errorf("could not parse headers from the HTTP response.")
	}
	client.m.Lock()
	defer client.m.Unlock()
	client.rateLimitingStats[api] = rateLimit{
		Limit:     resp.Header.Get("X-RateLimit-Limit"),
		Reset:     resp.Header.Get("X-RateLimit-Reset"),
		Period:    resp.Header.Get("X-RateLimit-Period"),
		Remaining: resp.Header.Get("X-RateLimit-Remaining"),
	}
	return nil
}

// GetRateLimitStats is a threadsafe getter to retrieve the rate limiting stats associated with the Client.
func (client *Client) GetRateLimitStats() map[string]rateLimit {
	client.m.Lock()
	defer client.m.Unlock()
	return client.rateLimitingStats
}
