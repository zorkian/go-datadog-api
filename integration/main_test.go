package integration

import (
	"github.com/zorkian/go-datadog-api"
	"log"
	"os"
	"testing"
)

var (
	apiKey string
	appKey string
	client *datadog.Client
)

func TestMain(m *testing.M) {
	apiKey = os.Getenv("DATADOG_API_KEY")
	appKey = os.Getenv("DATADOG_APP_KEY")

	if apiKey == "" || appKey == "" {
		log.Fatal("Please make sure to set the env variables 'DATADOG_API_KEY' and 'DATADOG_APP_KEY' before running this test")
	}

	client = datadog.NewClient(apiKey, appKey)
	os.Exit(m.Run())
}
