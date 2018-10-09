package datadog_test

import (
	"testing"

	"encoding/json"
	"github.com/stretchr/testify/assert"
	dd "github.com/zorkian/go-datadog-api"
)

func TestMonitorSerialization(t *testing.T) {

	raw := `
	{
		"id": 91879,
		"message": "We may need to add web hosts if this is consistently high.",
		"name": "Bytes received on host0",
		"options": {
			"no_data_timeframe": 20,
			"notify_audit": false,
			"notify_no_data": false,
			"silenced": {}
		},
		"org_id": 1499,
		"query": "avg(last_1h):sum:system.net.bytes_rcvd{host:host0} > 100",
		"type": "metric alert",
		"multi": false,
		"created": "2015-12-18T16:34:14.014039+00:00",
		"modified": "2015-12-18T16:34:14.014039+00:00",
		"state": {
			"groups": {
				"host:host0": {
					"last_nodata_ts": null,
					"last_notified_ts": 1481909160,
					"last_resolved_ts": 1481908200,
					"last_triggered_ts": 1481909160,
					"name": "host:host0",
					"status": "Alert",
					"triggering_value": {
						"from_ts": 1481909037,
						"to_ts": 1481909097,
						"value": 1000
					}
				}
			}
		}
	}`

	var monitor dd.Monitor
	err := json.Unmarshal([]byte(raw), &monitor)

	assert.Equal(t, err, nil)

	assert.Equal(t, *monitor.Id, 91879)
	assert.Equal(t, *monitor.State.Groups["host:host0"].Name, "host:host0")

}
