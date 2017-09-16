package datadog

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func TestUriForApi(t *testing.T) {
	c := Client{
		apiKey: "sample_api_key",
		appKey: "sample_app_key",
		HttpClient: &http.Client{},
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