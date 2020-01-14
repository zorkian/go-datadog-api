package datadog

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

var needKeysInQueryParams = []string{"/v1/series", "/v1/check_run", "/v1/events", "/v1/screen"}

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
	t.Run("Test all endpoints that need keys in query params", func(t *testing.T) {
		for _, api := range needKeysInQueryParams {
			uri, err := c.uriForAPI(api)
			assert.Nil(t, err)
			parsed, err := url.Parse(uri)
			assert.Nil(t, err)
			assert.Equal(t, parsed.Query().Get("api_key"), "sample_api_key")
			assert.Equal(t, parsed.Query().Get("application_key"), "sample_app_key")
		}
	})
	t.Run("Test an endpoint that doesn't need keys in query params", func(t *testing.T) {
		uri, err := c.uriForAPI("/v1/dashboard")
		assert.Nil(t, err)
		assert.Equal(t, "https://base.datadoghq.com/api/v1/dashboard", uri)
	})
}

func TestCreateRequest(t *testing.T) {
	c := Client{
		apiKey:       "sample_api_key",
		appKey:       "sample_app_key",
		baseUrl:      "https://base.datadoghq.com",
		HttpClient:   &http.Client{},
		RetryTimeout: 1000,
	}
	t.Run("Test an endpoint that doesn't need keys in query params", func(t *testing.T) {
		req, err := c.createRequest("GET", "/v1/dashboard", nil)
		assert.Nil(t, err)
		assert.Equal(t, "sample_api_key", req.Header.Get("DD-API-KEY"))
		assert.Equal(t, "sample_app_key", req.Header.Get("DD-APPLICATION-KEY"))
	})
	t.Run("Test endpoints that need keys in query params", func(t *testing.T) {
		for _, api := range needKeysInQueryParams {
			req, err := c.createRequest("GET", api, nil)
			assert.Nil(t, err)
			// we make sure that we *don't* have keys in query params, because some endpoints
			// fail if we send keys both in headers and query params
			assert.Equal(t, "", req.Header.Get("DD-API-KEY"))
			assert.Equal(t, "", req.Header.Get("DD-APPLICATION-KEY"))
		}
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
