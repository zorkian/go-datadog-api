package datadog

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUriForApi(t *testing.T) {
	c := Client{
		apiKey:       "sample_api_key",
		appKey:       "sample_app_key",
		baseUrl:      "https://app.datadoghq.com",
		HttpClient:   &http.Client{},
		RetryTimeout: 1000,
	}
	t.Run("Get Uri for api string with query string ", func(t *testing.T) {
		uri, err := c.uriForAPI("/v1/events?type=critical")
		assert.Nil(t, err)
		assert.Equal(t, "https://app.datadoghq.com/api/v1/events?api_key=sample_api_key&application_key=sample_app_key&type=critical", uri)

	})
	t.Run("Get Uri for api without query string", func(t *testing.T) {
		uri, err := c.uriForAPI("/v1/events")
		assert.Nil(t, err)
		assert.Equal(t, "https://app.datadoghq.com/api/v1/events?api_key=sample_api_key&application_key=sample_app_key", uri)
	})
}

func TestRedactError(t *testing.T) {
	c := Client{
		apiKey:       "sample_api_key",
		appKey:       "sample_app_key",
		baseUrl:      "https://app.datadoghq.com",
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
