// +build integration

package datadog

import (
	"github.com/zorkian/go-datadog-api"
	"log"
	"os"
)

// TODO: push this out to helper file before creating PR
var (
	apiKey string
	appKey string
	client *datadog.Client
)

func initTest() *datadog.Client {
	apiKey = os.Getenv("DATADOG_API_KEY")
	appKey = os.Getenv("DATADOG_APP_KEY")

	if apiKey == "" || appKey == "" {
		log.Fatal("Please make sure to set the env variables 'DATADOG_API_KEY' and 'DATADOG_APP_KEY' before running this test")
	}

	return datadog.NewClient(apiKey, appKey)
}
