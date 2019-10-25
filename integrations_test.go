package datadog

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

)

func TestIntegrationWebhookGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/webhooks_response.json")
		if err != nil {
			t.Fatal(err)
		}
		w.Write(response)
	}))

	defer ts.Close()

	client := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	integrationWebhook, err := client.GetIntegrationWebhook()
	assert.Nil(t, err)
	assert.Equal(t, expectedIntegrationWebhook, integrationWebhook)
}

var expectedIntegrationWebhook = &IntegrationWebhook{
	Hooks: []IntegrationWebhookHook{
		{
			Name:				String("Test"),
			Url:				String("http://example.com"),
			EncodeAsForm:		CBool(false),
			UseCustomPayload:	CBool(true),
			CustomPayload:		String("{\n\"body\": \"$EVENT_MSG\",\n    \"last_updated\": \"$LAST_UPDATED\",\n    \"event_type\": \"$EVENT_TYPE\",\n    \"title\": \"$EVENT_TITLE\",\n    \"date\": \"$DATE\",\n    \"org\": {\n        \"id\": \"$ORG_ID\",\n        \"name\": \"$ORG_NAME\"\n    },\n    \"id\": \"$ID\"\n}"),
			Headers:			String("X-Dummy-Header: Dummy-Value"),
		},
	},
}
