package datadog_test

import (
	"testing"

	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
)

func TestMonitorOptionsEvaluationDelayUnmarshal(t *testing.T) {
	type TestCase struct {
		jsonData      []byte
		expectedValue int
	}

	testCases := []TestCase{
		{[]byte("{\"evaluation_delay\": 0}"), 0},
		{[]byte("{\"evaluation_delay\": \"0\"}"), 0},
		{[]byte("{\"evaluation_delay\": \"\"}"), 0},
		{[]byte("{\"evaluation_delay\": 10}"), 10},
		{[]byte("{\"evaluation_delay\": \"10\"}"), 10},
	}

	for _, testCase := range testCases {
		options := datadog.Options{}

		if err := json.Unmarshal(testCase.jsonData, &options); err != nil {
			t.Fatalf("Unmarshalling evaluation_delay failed when it shouldn't: %s", err)
		}

		assert.Equal(t, testCase.expectedValue, int(options.GetEvaluationDelay()))
	}
}
