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

func TestMonitorSerializationForLogAlert(t *testing.T) {

	raw := `
	{
		"tags": [
		  "app:webserver",
		  "frontend"
		],
		"deleted": null,
		"query": "logs(\"env:develop\").index(\"main\").rollup(\"count\").last(\"5m\") > 500",
		"message": "Monitor Log Count",
		"id": 91879,
		"multi": false,
		"name": "Monitor Log Count",
		"created": "2018-12-06T08:26:17.235509+00:00",
		"created_at": 1544084777000,
		"org_id": 1499,
		"modified": "2018-12-06T08:26:17.235509+00:00",
		"overall_state_modified": null,
		"overall_state": "No Data",
		"type": "log alert",
		"options": {
		  "notify_audit": false,
		  "locked": false,
		  "timeout_h": 0,
		  "silenced": {},
		  "enable_logs_sample": true,
		  "thresholds": {
			"comparison": ">",
			"critical": 1000,
			"period": {
			  "seconds": 300,
			  "text": "5 minutes",
			  "value": "last_5m",
			  "name": "5 minute average",
			  "unit": "minutes"
			},
			"timeAggregator": "avg"
		  },
		  "queryConfig": {
			"logset": {
			  "id": "1499",
			  "name": "main"
			},
			"timeRange": {
			  "to": 1539675566736,
			  "live": true,
			  "from": 1539661166736
			},
			"queryString": "env:develop",
			"queryIsFailed": false
		  },
		  "new_host_delay": 300,
		  "notify_no_data": false,
		  "renotify_interval": 0,
		  "evaluation_delay": 500,
		  "no_data_timeframe": 2
		}
	  }
	  `

	var monitor dd.Monitor
	err := json.Unmarshal([]byte(raw), &monitor)

	assert.Equal(t, err, nil)

	assert.Equal(t, 91879, *monitor.Id)
	assert.Equal(t, true, *monitor.Options.EnableLogsSample)
	assert.Equal(t, json.Number("1499"), *monitor.Options.QueryConfig.LogSet.ID)
	assert.Equal(t, json.Number("1539661166736"), *monitor.Options.QueryConfig.TimeRange.From)
	assert.Equal(t, "env:develop", *monitor.Options.QueryConfig.QueryString)
}
