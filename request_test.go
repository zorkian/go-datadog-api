package datadog

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUriForApi(t *testing.T) {
	c := Client{
		apiKey:       "sample_api_key",
		appKey:       "sample_app_key",
		baseUrl:      "https://base.datadoghq.com",
		HttpClient:   &http.Client{},
		RetryTimeout: 1000,
	}
	t.Run("Get Uri for api string with query string", func(t *testing.T) {
		uri, err := c.uriForAPI("/v1/events?type=critical")
		assert.Nil(t, err)
		assert.Equal(t, "https://base.datadoghq.com/api/v1/events?api_key=sample_api_key&application_key=sample_app_key&type=critical", uri)

	})
	t.Run("Get Uri for api without query string", func(t *testing.T) {
		uri, err := c.uriForAPI("/v1/events")
		assert.Nil(t, err)
		assert.Equal(t, "https://base.datadoghq.com/api/v1/events?api_key=sample_api_key&application_key=sample_app_key", uri)
	})
}

func TestRedactError(t *testing.T) {
	c := Client{
		apiKey:       "sample_api_key",
		appKey:       "sample_app_key",
		baseUrl:      "https://base.datadoghq.com",
		HttpClient:   &http.Client{},
		RetryTimeout: 1000,
	}
	t.Run("Error containing api key in string is correctly redacted", func(t *testing.T) {
		var leakErr = fmt.Errorf("Error test: %s,%s", c.apiKey, c.apiKey)
		var redactedErr = c.redactError(leakErr)

		if assert.NotNil(t, redactedErr) {
			assert.Equal(t, "Error test: redacted,redacted", redactedErr.Error())
		}
	})
	t.Run("Error containing application key in string is correctly redacted", func(t *testing.T) {
		var leakErr = fmt.Errorf("Error test: %s,%s", c.appKey, c.appKey)
		var redactedErr = c.redactError(leakErr)

		if assert.NotNil(t, redactedErr) {
			assert.Equal(t, "Error test: redacted,redacted", redactedErr.Error())
		}
	})
	t.Run("Nil error returns nil", func(t *testing.T) {
		var harmlessErr error = nil
		var redactedErr = c.redactError(harmlessErr)

		assert.Nil(t, redactedErr)
	})
}

func makeTestServer(code int, response string) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/something", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write([]byte(response))
	})
	server := httptest.NewServer(mux)
	return server
}

func TestErrorHandling(t *testing.T) {
	c := Client{
		apiKey:       "sample_api_key",
		appKey:       "sample_app_key",
		HttpClient:   &http.Client{},
		RetryTimeout: 1000,
	}
	for _, code := range []int{401, 403, 500, 502} {
		t.Run(fmt.Sprintf("Returns error on http code %d", code), func(t *testing.T) {
			s := makeTestServer(code, "")
			defer s.Close()
			c.SetBaseUrl(s.URL)

			for _, method := range []string{"GET", "POST", "PUT"} {
				err := c.doJsonRequest(method, "/v1/something", nil, nil)
				assert.NotNil(t, err)
			}
		})
	}
	t.Run("Returns error if status is error", func(t *testing.T) {
		s := makeTestServer(200, `{"status": "error", "error": "something wrong"}`)
		defer s.Close()
		c.SetBaseUrl(s.URL)

		for _, method := range []string{"GET", "POST", "PUT"} {
			err := c.doJsonRequest(method, "/v1/something", nil, nil)
			if assert.NotNil(t, err) {
				assert.Contains(t, err.Error(), "something wrong")
			}
		}
	})
	t.Run("Does not return error if status is ok", func(t *testing.T) {
		s := makeTestServer(200, `{"status": "ok"}`)
		defer s.Close()
		c.SetBaseUrl(s.URL)

		for _, method := range []string{"GET", "POST", "PUT"} {
			err := c.doJsonRequest(method, "/v1/something", nil, nil)
			assert.Nil(t, err)
		}
	})
}
