package integration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
)

func TestInvalidAuth(t *testing.T) {
	// Override the correct credentials
	c := datadog.NewClient("INVALID", "INVALID")

	valid, err := c.Validate()
	if err != nil {
		t.Fatalf("Testing authentication failed when it shouldn't: %s", err)
	}

	assert.Equal(t, valid, false)
}

func TestValidAuth(t *testing.T) {
	valid, err := client.Validate()

	if err != nil {
		t.Fatalf("Testing authentication failed when it shouldn't: %s", err)
	}

	assert.Equal(t, valid, true)
}

func TestBaseUrl(t *testing.T) {
	t.Run("Base url defaults to https://app.datadoghq.com", func(t *testing.T) {
		assert.Empty(t, os.Getenv("DATADOG_HOST"))

		c := datadog.NewClient("abc", "def")
		assert.Equal(t, "https://app.datadoghq.com", c.GetBaseUrl())
	})

	t.Run("Base url defaults DATADOG_HOST environment variable if set", func(t *testing.T) {
		os.Setenv("DATADOG_HOST", "https://custom.datadoghq.com")
		defer os.Unsetenv("DATADOG_HOST")

		c := datadog.NewClient("abc", "def")
		assert.Equal(t, "https://custom.datadoghq.com", c.GetBaseUrl())
	})

	t.Run("Base url can be set through the attribute setter", func(t *testing.T) {
		c := datadog.NewClient("abc", "def")
		c.SetBaseUrl("https://another.datadoghq.com")
		assert.Equal(t, "https://another.datadoghq.com", c.GetBaseUrl())
	})
}
