/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2016 by authors and contributors.
 */

package datadog_test

import (
	"testing"

	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
)

func TestBool(t *testing.T) {

}

func TestHelperGetBoolSet(t *testing.T) {
	// Assert that we were able to get the string from a pointer field
	m := getTestMonitor()

	if attr, ok := datadog.GetBool(m.Options.NotifyNoData); ok {
		assert.Equal(t, true, attr)
	}
}

func TestHelperGetBoolNotSet(t *testing.T) {
	// Assert that we were able to get the string from a pointer field
	m := getTestMonitor()

	_, ok := datadog.GetBool(m.Options.NotifyAudit)
	assert.Equal(t, false, ok)
}

func TestHelperStringSet(t *testing.T) {
	// Assert that we were able to get the string from a pointer field
	m := getTestMonitor()

	if attr, ok := datadog.GetString(m.Name); ok {
		assert.Equal(t, "Test monitor", attr)
	}
}

func TestHelperStringNotSet(t *testing.T) {
	// Assert GetString returned false for an unset value
	m := getTestMonitor()

	_, ok := datadog.GetString(m.Message)
	assert.Equal(t, false, ok)
}

func TestHelperIntSet(t *testing.T) {
	// Assert that we were able to get the string from a pointer field
	m := getTestMonitor()

	if attr, ok := datadog.GetInt(m.Id); ok {
		assert.Equal(t, 1, attr)
	}
}

func TestHelperIntNotSet(t *testing.T) {
	// Assert GetString returned false for an unset value
	m := getTestMonitor()

	_, ok := datadog.GetInt(m.Options.RenotifyInterval)
	assert.Equal(t, false, ok)
}

func TestHelperGetJsonNumberSet(t *testing.T) {
	// Assert that we were able to get a JSON Number from a pointer field
	m := getTestMonitor()

	if attr, ok := datadog.GetJsonNumber(m.Options.Thresholds.Ok); ok {
		assert.Equal(t, json.Number(2), attr)
	}
}

func TestHelperGetJsonNumberNotSet(t *testing.T) {
	// Assert GetString returned false for an unset value
	m := getTestMonitor()

	_, ok := datadog.GetJsonNumber(m.Options.Thresholds.Warning)

	assert.Equal(t, false, ok)
}

func getTestMonitor() *datadog.Monitor {

	o := &datadog.Options{
		NotifyNoData:    datadog.Bool(true),
		Locked:          datadog.Bool(false),
		NoDataTimeframe: 60,
		Silenced:        map[string]int{},
		Thresholds: &datadog.ThresholdCount{
			Ok: datadog.JsonNumber(json.Number(2)),
		},
	}

	return &datadog.Monitor{
		Query:   datadog.String("avg(last_15m):avg:system.disk.in_use{*} by {host,device} > 0.8"),
		Name:    datadog.String("Test monitor"),
		Id:      datadog.Int(1),
		Options: o,
		Type:    datadog.String("metric alert"),
		Tags:    make([]string, 0),
	}
}
